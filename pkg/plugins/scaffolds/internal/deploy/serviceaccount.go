package deploy

import (
	"csibuilder/pkg/machinery"
	"path/filepath"
)

var _ machinery.Template = &ServiceAccountYaml{}

// ServiceAccountYaml scaffolds a file that defines ServiceAccount yaml for csi node & controller
type ServiceAccountYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *ServiceAccountYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "deploy/serviceaccount.yaml")
	}

	f.TemplateBody = serviceaccountTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const serviceaccountTemplate = `apiVersion: v1
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
`
