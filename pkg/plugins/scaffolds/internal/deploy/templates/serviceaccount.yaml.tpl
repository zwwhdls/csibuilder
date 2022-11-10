apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-controller
  namespace: default
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-node
  namespace: default