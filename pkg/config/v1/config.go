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

package v1

import (
	"gopkg.in/yaml.v2"

	"csibuilder/pkg/config"
)

// VersionV1 is the config.Version for project configuration 1
var VersionV1 = config.Version{Number: 1}

type Config struct {
	// Version
	Version config.Version `json:"version"`

	Name string
	Repo string
}

var _ config.Config = &Config{}

// New returns a new config.Config
func New() config.Config {
	return &Config{Version: VersionV1}
}

func init() {
	config.Register(VersionV1, New)
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

func (c *Config) GetVersion() config.Version {
	return c.Version
}

func (c *Config) GetProjectName() string {
	return c.Name
}

func (c *Config) SetProjectName(name string) error {
	c.Name = name
	return nil
}

func (c *Config) MarshalYAML() ([]byte, error) {
	content, err := yaml.Marshal(c)
	if err != nil {
		return nil, config.MarshalError{Err: err}
	}

	return content, nil
}

func (c *Config) UnmarshalYAML(b []byte) error {
	if err := yaml.UnmarshalStrict(b, c); err != nil {
		return config.UnmarshalError{Err: err}
	}
	return nil
}
