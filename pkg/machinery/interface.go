package machinery

import (
	"csibuilder/pkg/model"
	"text/template"
)

type Builder interface {
	// GetPath returns the path to the file location
	GetPath() string
	// GetIfExistsAction returns the behavior when creating a file that already exists
	GetIfExistsAction() IfExistsAction
}

// Inserter is a file builder that inserts code fragments in marked positions
type Inserter interface {
	Builder
	// GetMarkers returns the different markers where code fragments will be inserted
	GetMarkers() []Marker
	// GetCodeFragments returns a map that binds markers to code fragments
	GetCodeFragments() CodeFragmentsMap
}

// Template is file builder based on a file template
type Template interface {
	Builder
	// GetBody returns the template body
	GetBody() string
	// SetTemplateDefaults sets the default values for templates
	SetTemplateDefaults() error
}

// UseCustomFuncMap allows a template to use a custom template.FuncMap instead of the default FuncMap.
type UseCustomFuncMap interface {
	// GetFuncMap returns a custom FuncMap.
	GetFuncMap() template.FuncMap
}

// HasRepository allows the repository to be used on a template
type HasRepository interface {
	// InjectRepository sets the template repository
	InjectRepository(string)
}

// HasResource allows a resource to be used on a template
type HasResource interface {
	// InjectResource sets the template resource
	InjectResource(*model.Resource)
}
