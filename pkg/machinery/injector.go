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

package machinery

import (
	"csibuilder/pkg/config"
	"csibuilder/pkg/model"
)

// injector is used to inject certain fields to file templates.
type injector struct {
	// config is the project configuration
	config config.Config

	// resource contains the information of the CSI that is being scaffolded.
	resource *model.Resource

	// boilerplate is the copyright comment added at the top of scaffolded files.
	boilerplate string
}

// injectInto injects fields from the universe into the builder
func (i injector) injectInto(builder Builder) {
	// Inject project configuration
	if i.config != nil {
		if builderWithRepository, hasRepository := builder.(HasRepository); hasRepository {
			builderWithRepository.InjectRepository(i.config.GetRepository())
		}
		if builderWithVersion, hasVersion := builder.(HasVersion); hasVersion {
			builderWithVersion.InjectGoVersion(i.config.GetGoVersion())
		}
	}
	// Inject resource
	if i.resource != nil {
		if builderWithResource, hasResource := builder.(HasResource); hasResource {
			builderWithResource.InjectResource(i.resource)
		}
	}
	// Inject boilerplate
	if builderWithBoilerplate, hasBoilerplate := builder.(HasBoilerplate); hasBoilerplate {
		builderWithBoilerplate.InjectBoilerplate(i.boilerplate)
	}
}
