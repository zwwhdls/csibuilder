apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: {{ .Resource.CSIName }}
spec:
  attachRequired: {{ .Resource.AttachRequired }}
  podInfoOnMount: false