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

package main

import (
	"csibuilder/pkg/model"
	"github.com/spf13/pflag"
)

type resourceOptions struct {
	CSIName        string
	AttachRequired bool
}

func bindResourceFlags(fs *pflag.FlagSet) *resourceOptions {
	options := &resourceOptions{}

	fs.StringVar(&options.CSIName, "csi", "", "csi name")
	fs.BoolVar(&options.AttachRequired, "attach", false, "attach required in csi")

	return options
}

func (opts resourceOptions) newResource() *model.Resource {
	return &model.Resource{
		CSIName:        opts.CSIName,
		AttachRequired: opts.AttachRequired,
	}
}
