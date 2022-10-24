package internal

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"os"
	"path/filepath"
)

var _ machinery.Template = &GoMod{}

// GoMod scaffolds a file that defines the project dependencies
type GoMod struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *GoMod) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "go.mod")
	}

	if f.TemplatePath == "" {
		return fmt.Errorf("can not get template path")
	}

	templateFile := filepath.Join(f.TemplatePath, "go.mod.tpl")
	body, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
