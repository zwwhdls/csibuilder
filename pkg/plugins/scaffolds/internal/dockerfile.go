package internal

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"path/filepath"
)

var _ machinery.Template = &Dockerfile{}

// Dockerfile scaffolds Dockerfile file
type Dockerfile struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *Dockerfile) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "Dockerfile")
	}
	if f.TemplatePath == "" {
		return fmt.Errorf("can not get template path")
	}

	//templateFile := filepath.Join(f.TemplatePath, "Dockerfile.tpl")
	body, err := tplFS.ReadFile("templates/Dockerfile.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
