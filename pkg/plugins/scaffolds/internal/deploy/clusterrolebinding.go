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
)

var _ machinery.Template = &ClusterRoleBindingYaml{}

// ClusterRoleBindingYaml scaffolds a file that defines ClusterRoleBinding yaml for csi node & controller
type ClusterRoleBindingYaml struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	Force bool
}

// SetTemplateDefaults implements file.Template
func (f *ClusterRoleBindingYaml) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "deploy/clusterrolebinding.yaml"
	}

	body, err := tplFS.ReadFile("templates/clusterrolebinding.yaml.tpl")
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}
