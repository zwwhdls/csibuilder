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
	"embed"

	"csibuilder/pkg/machinery"
)

var _ machinery.Template = &ClusterRoleYaml{}

// ClusterRoleYaml scaffolds a file that defines ClusterRole yaml for csi node & controller
type ClusterRoleYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin
	machinery.BoilerplateMixin

	Force bool
}

//go:embed templates/*
var tplFS embed.FS

// SetTemplateDefaults implements file.Template
func (f *ClusterRoleYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "deploy/clusterrole.yaml"
	}

	body, err := tplFS.ReadFile("templates/clusterrole.yaml.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
