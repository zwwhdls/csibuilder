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

package model

type Config struct {
	Repo         string
	TemplatePath string
}

// GetRepository returns the project repository.
func (c *Config) GetRepository() string {
	return c.Repo
}

// SetRepository sets the project repository.
func (c *Config) SetRepository(repository string) error {
	c.Repo = repository
	return nil
}

// GetTemplatePath returns the template path
func (c *Config) GetTemplatePath() string {
	return c.TemplatePath
}

// SetTemplatePath sets the template path.
func (c *Config) SetTemplatePath(templatePath string) error {
	c.TemplatePath = templatePath
	return nil
}
