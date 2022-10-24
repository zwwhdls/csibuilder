package deploy

import (
	"csibuilder/pkg/machinery"
	"path/filepath"
)

var _ machinery.Template = &ClusterRoleBindingYaml{}

// ClusterRoleBindingYaml scaffolds a file that defines ClusterRoleBinding yaml for csi node & controller
type ClusterRoleBindingYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *ClusterRoleBindingYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "deploy/clusterrolebinding.yaml")
	}

	f.TemplateBody = clusterrolebindingTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const clusterrolebindingTemplate = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: csi-node
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: csi-node
subjects:
  - kind: ServiceAccount
    name: csi-node
    namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: csi-provisioner
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: csi-provisioner
subjects:
  - kind: ServiceAccount
    name: csi-controller
    namespace: default
`
