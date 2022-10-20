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

// DomainMixin provides templates with a injectable domain field
type DomainMixin struct {
	// Domain is the domain for the APIs
	Domain string
}

// InjectDomain implements HasDomain
func (m *DomainMixin) InjectDomain(domain string) {
	if m.Domain == "" {
		m.Domain = domain
	}
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

// MultiGroupMixin provides templates with a injectable multi-group flag field
type MultiGroupMixin struct {
	// MultiGroup is the multi-group flag
	MultiGroup bool
}

// InjectMultiGroup implements HasMultiGroup
func (m *MultiGroupMixin) InjectMultiGroup(flag bool) {
	m.MultiGroup = flag
}

// ComponentConfigMixin provides templates with a injectable component-config flag field
type ComponentConfigMixin struct {
	// ComponentConfig is the component-config flag
	ComponentConfig bool
}

// InjectComponentConfig implements HasComponentConfig
func (m *ComponentConfigMixin) InjectComponentConfig(flag bool) {
	m.ComponentConfig = flag
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
