package internal

import (
	"csibuilder/pkg/machinery"
	"path/filepath"
)

var _ machinery.Template = &GoMod{}

// GoMod scaffolds a file that defines the project dependencies
type GoMod struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin

	ControllerRuntimeVersion string
}

// SetTemplateDefaults implements file.Template
func (f *GoMod) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "go.mod")
	}

	f.TemplateBody = goModTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const goModTemplate = `
module {{ .Repo }}

go 1.18

require (
	github.com/container-storage-interface/spec v1.6.0
	k8s.io/klog v1.0.0
)

`
