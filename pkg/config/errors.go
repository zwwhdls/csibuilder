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

import (
	"fmt"
)

// UnsupportedVersionError is returned by New when a project configuration version is not supported.
type UnsupportedVersionError struct {
	Version Version
}

// Error implements error interface
func (e UnsupportedVersionError) Error() string {
	return fmt.Sprintf("version %s is not supported", e.Version)
}

// LoadError wraps errors yielded by Store.Load and Store.LoadFrom methods
type LoadError struct {
	Err error
}

// MarshalError is returned by Config.Marshal when something went wrong while marshalling to YAML
type MarshalError struct {
	Err error
}

// Error implements error interface
func (e MarshalError) Error() string {
	return fmt.Sprintf("error marshalling project configuration: %v", e.Err)
}

// Unwrap implements Wrapper interface
func (e MarshalError) Unwrap() error {
	return e.Err
}

// UnmarshalError is returned by Config.Unmarshal when something went wrong while unmarshalling from YAML
type UnmarshalError struct {
	Err error
}

// Error implements error interface
func (e UnmarshalError) Error() string {
	return fmt.Sprintf("error unmarshalling project configuration: %v", e.Err)
}

// Unwrap implements Wrapper interface
func (e UnmarshalError) Unwrap() error {
	return e.Err
}

// Error implements error interface
func (e LoadError) Error() string {
	return fmt.Sprintf("unable to load the configuration: %v", e.Err)
}

// Unwrap implements Wrapper interface
func (e LoadError) Unwrap() error {
	return e.Err
}

// SaveError wraps errors yielded by Store.Save and Store.SaveTo methods
type SaveError struct {
	Err error
}

// Error implements error interface
func (e SaveError) Error() string {
	return fmt.Sprintf("unable to save the configuration: %v", e.Err)
}

// Unwrap implements Wrapper interface
func (e SaveError) Unwrap() error {
	return e.Err
}
