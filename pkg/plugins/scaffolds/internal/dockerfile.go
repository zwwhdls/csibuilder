package internal

import (
	"csibuilder/pkg/machinery"
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

	f.TemplateBody = dockerfileTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const dockerfileTemplate = `
FROM golang:1.18-buster

ARG GOPROXY

WORKDIR /workspace
COPY . .
ENV GOPROXY=${GOPROXY:-https://proxy.golang.org}

RUN make csi
RUN chmod u+x /workspace/bin/csi

ENTRYPOINT ["/workspace/bin/csi"]

`
