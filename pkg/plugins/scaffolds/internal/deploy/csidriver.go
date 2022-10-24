package deploy

import (
	"csibuilder/pkg/machinery"
	"path/filepath"
)

var _ machinery.Template = &CSIDriverYaml{}

// CSIDriverYaml scaffolds a file that defines CSIDriver yaml for csi
type CSIDriverYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *CSIDriverYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "deploy/csidriver.yaml")
	}

	f.TemplateBody = csidriverTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const csidriverTemplate = `apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: {{ .Resource.CSIName }}
spec:
  attachRequired: {{ .Resource.AttachRequired }}
  podInfoOnMount: false
`
