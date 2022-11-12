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

package config

// Config defines the interface that project configuration types must follow.
type Config interface {
	/* Version */

	// GetVersion returns the current project version.
	GetVersion() Version

	/* String fields */

	// SetGoVersion sets the project Go version.
	SetGoVersion(goVersion string) error

	// GetGoVersion returns the project Go version.
	GetGoVersion() string

	// GetRepository returns the project repository.
	GetRepository() string
	// SetRepository sets the project repository.
	SetRepository(repository string) error

	// GetProjectName returns the project name.
	// This method was introduced in project version 3.
	GetProjectName() string
	// SetProjectName sets the project name.
	// This method was introduced in project version 3.
	SetProjectName(name string) error

	/* Persistence */

	// MarshalYAML returns the YAML representation of the Config.
	MarshalYAML() ([]byte, error)
	// UnmarshalYAML loads the Config fields from its YAML representation.
	UnmarshalYAML([]byte) error
}
