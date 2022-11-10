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

package deploy

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"path/filepath"
)

var _ machinery.Template = &StatefulSetYaml{}

// StatefulSetYaml scaffolds a file that defines csi controller deploy yaml
type StatefulSetYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *StatefulSetYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "deploy/statefulset.yaml")
	}
	fmt.Println(f.Path)

	if f.TemplatePath == "" {
		return fmt.Errorf("can not get template path")
	}

	body, err := tplFS.ReadFile("templates/statefulset.yaml.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
