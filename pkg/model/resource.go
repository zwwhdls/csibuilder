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

import "strings"

// Resource contains the information required to scaffold files for a resource.
type Resource struct {
	// Path is the path to the go package where the types are defined.
	Path           string `json:"path,omitempty"`
	CSIName        string `json:"csiName"`
	AttachRequired bool   `json:"attachRequired,omitempty"`
}

func (r Resource) Replacer() *strings.Replacer {
	var replacements []string

	return strings.NewReplacer(replacements...)
}
