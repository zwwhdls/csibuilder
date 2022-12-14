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

package internal

import (
	"embed"

	"csibuilder/pkg/machinery"
)

var _ machinery.Template = &Main{}

// Main scaffolds the main.go
// nolint:maligned
type Main struct {
	machinery.TemplateMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin
	machinery.BoilerplateMixin

	Force bool
}

//go:embed templates/*
var tplFS embed.FS

func (f *Main) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "main.go"
	}
	f.Path = f.Resource.Replacer().Replace(f.Path)

	body, err := tplFS.ReadFile("templates/main.go.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	if f.Force {
		f.IfExistsAction = machinery.OverwriteFile
	} else {
		f.IfExistsAction = machinery.Error
	}
	return nil
}
