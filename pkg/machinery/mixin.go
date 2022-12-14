/*
 Copyright 2022 CSIBuilder

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License
*/

package machinery

import "csibuilder/pkg/model"

// PathMixin provides file builders with a path field
type PathMixin struct {
	// Path is the of the file
	Path string
}

// GetPath implements Builder
func (t *PathMixin) GetPath() string {
	return t.Path
}

// TemplatePathMixin provides file builders with a path field
type TemplatePathMixin struct {
	// TemplatePath is the path of the template file
	TemplatePath string
}

// InjectTemplatePath implements HasTemplatePath
func (t *TemplatePathMixin) InjectTemplatePath(templatePath string) {
	t.TemplatePath = templatePath
}

// IfExistsActionMixin provides file builders with a if-exists-action field
type IfExistsActionMixin struct {
	// IfExistsAction determines what to do if the file exists
	IfExistsAction IfExistsAction
}

// GetIfExistsAction implements Builder
func (t *IfExistsActionMixin) GetIfExistsAction() IfExistsAction {
	return t.IfExistsAction
}

// TemplateMixin is the mixin that should be embedded in Template builders
type TemplateMixin struct {
	PathMixin
	TemplatePathMixin
	IfExistsActionMixin

	// TemplateBody is the template body to execute
	TemplateBody string
}

// GetBody implements Template
func (t *TemplateMixin) GetBody() string {
	return t.TemplateBody
}

// InserterMixin is the mixin that should be embedded in Inserter builders
type InserterMixin struct {
	PathMixin
}

// GetIfExistsAction implements Builder
func (t *InserterMixin) GetIfExistsAction() IfExistsAction {
	// Inserter builders always need to overwrite previous files
	return OverwriteFile
}

// RepositoryMixin provides templates with a injectable repository field
type RepositoryMixin struct {
	// Repo is the go project package path
	Repo string
}

// InjectRepository implements HasRepository
func (m *RepositoryMixin) InjectRepository(repository string) {
	if m.Repo == "" {
		m.Repo = repository
	}
}

// GoVersionMixin provides templates with a injectable go version field
type VersionMixin struct {
	GoVersion string
	Number    int
	Stage     model.Stage
}

// InjectGoVersion implements HasGoVersion
func (m *VersionMixin) InjectGoVersion(goVersion string) {
	if m.GoVersion == "" {
		m.GoVersion = goVersion
	}
}

// ProjectNameMixin provides templates with an injectable project name field.
type ProjectNameMixin struct {
	ProjectName string
}

// InjectProjectName implements HasProjectName.
func (m *ProjectNameMixin) InjectProjectName(projectName string) {
	if m.ProjectName == "" {
		m.ProjectName = projectName
	}
}

// ResourceMixin provides templates with a injectable resource field
type ResourceMixin struct {
	Resource *model.Resource
}

// InjectResource implements HasResource
func (m *ResourceMixin) InjectResource(res *model.Resource) {
	if m.Resource == nil {
		m.Resource = res
	}
}

// BoilerplateMixin provides templates with a injectable boilerplate field
type BoilerplateMixin struct {
	// Boilerplate is the contents of a Boilerplate go header file
	Boilerplate string
}

// InjectBoilerplate implements HasBoilerplate
func (m *BoilerplateMixin) InjectBoilerplate(boilerplate string) {
	if m.Boilerplate == "" {
		m.Boilerplate = boilerplate
	}
}
