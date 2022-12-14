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

package csi

import (
	"path/filepath"

	"csibuilder/pkg/machinery"
)

var _ machinery.Template = &Driver{}

// Driver scaffolds the file that defines csi driver
// nolint:maligned
type Driver struct {
	machinery.TemplateMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin
	machinery.BoilerplateMixin

	Force bool
}

func (c *Driver) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join("pkg/csi", "driver.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)

	body, err := tplFS.ReadFile("templates/driver.go.tpl")
	if err != nil {
		return err
	}
	c.TemplateBody = string(body)

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
	}
	return nil
}
