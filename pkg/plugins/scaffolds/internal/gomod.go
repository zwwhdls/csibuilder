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
	"csibuilder/pkg/machinery"
	"fmt"
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
	fmt.Println(f.Path)

	if f.TemplatePath == "" {
		return fmt.Errorf("can not get template path")
	}

	body, err := tplFS.ReadFile("templates/go.mod.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
