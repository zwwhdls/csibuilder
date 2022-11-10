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
}

// HasTemplatePath is file builder based on a file template
type HasTemplatePath interface {
	// InjectTemplatePath sets the template path for templates
	InjectTemplatePath(templatePath string)
}

// Template is file builder based on a file template
type Template interface {
	Builder
	HasTemplatePath
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

// HasBoilerplate allows a boilerplate to be used on a template
type HasBoilerplate interface {
	// InjectBoilerplate sets the template boilerplate
	InjectBoilerplate(string)
}
