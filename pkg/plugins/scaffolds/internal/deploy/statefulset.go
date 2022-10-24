package deploy

import (
	"csibuilder/pkg/machinery"
	"path/filepath"
)

var _ machinery.Template = &StatefulSetYaml{}

// StatefulSetYaml scaffolds a file that defines csi controller deploy yaml
type StatefulSetYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *StatefulSetYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "deploy/statefulset.yaml")
	}

	f.TemplateBody = statefulsetTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const statefulsetTemplate = `apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app.kubernetes.io/component: controller
    app.kubernetes.io/name: {{ .Resource.CSIName }}-controller
  name: {{ .Resource.CSIName }}-controller
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Resource.CSIName }}-csi-controller
  serviceName: csi-controller
  template:
    metadata:
      labels:
        app: {{ .Resource.CSIName }}-csi-controller
    spec:
      priorityClassName: system-cluster-critical
      serviceAccountName: csi-controller
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
      containers:
        - args:
            - --endpoint=$(CSI_ENDPOINT)
            - --logtostderr
            - --nodeid=$(NODE_NAME)
          env:
            - name: CSI_ENDPOINT
              value: unix:///var/lib/csi/sockets/pluginproxy/csi.sock
            - name: NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          image: csi-image
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 3
          name: csi-plugin
          ports:
            - containerPort: 9909
              name: healthz
              protocol: TCP
          securityContext:
            capabilities:
              add:
                - SYS_ADMIN
            privileged: true
          volumeMounts:
            - mountPath: /var/lib/csi/sockets/pluginproxy/
              name: socket-dir
        - args:
            - --csi-address=$(ADDRESS)
            - --timeout=60s
            - --v=5
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          image: quay.io/k8scsi/csi-provisioner:v1.6.0
          name: csi-provisioner
          volumeMounts:
            - mountPath: /var/lib/csi/sockets/pluginproxy/
              name: socket-dir
        - args:
            - --csi-address=$(ADDRESS)
            - --health-port=$(HEALTH_PORT)
          env:
            - name: ADDRESS
              value: /csi/csi.sock
            - name: HEALTH_PORT
              value: "9909"
          image: quay.io/k8scsi/livenessprobe:v1.1.0
          name: liveness-probe
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
      volumes:
        - emptyDir: {}
          name: socket-dir
`
