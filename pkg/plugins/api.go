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

package plugins

import (
	"fmt"

	"github.com/spf13/pflag"

	"csibuilder/pkg/config"
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins/scaffolds"
	"csibuilder/pkg/util"
)

// DefaultMainPath is default file path of main.go
const DefaultMainPath = "main.go"

type CreateAPISubcommand struct {
	config   config.Config
	resource *model.Resource

	// Check if we have to scaffold resource and/or controller
	resourceFlag   *pflag.Flag
	controllerFlag *pflag.Flag
	repo           string

	// force indicates that the resource should be created even if it already exists
	force bool

	// runMake indicates whether to run make or not after scaffolding APIs
	runMake bool
}

func (p *CreateAPISubcommand) InjectResource(res *model.Resource) error {
	p.resource = res

	return nil
}

func (p *CreateAPISubcommand) InjectConfig(conf config.Config) error {
	p.config = conf

	// Try to guess repository if flag is not set.
	if conf.GetRepository() == "" {
		repoPath, err := util.FindCurrentRepo()
		if err != nil {
			return fmt.Errorf("error finding current repository: %v", err)
		}
		p.repo = repoPath
		conf.SetRepository(repoPath)
	}
	p.repo = conf.GetRepository()

	return nil
}

func (p *CreateAPISubcommand) PreScaffold(machinery.Filesystem) error {
	// check if main.go is present in the root directory
	//if _, err := os.Stat(DefaultMainPath); os.IsNotExist(err) {
	//	return fmt.Errorf("%s file should present in the root directory", DefaultMainPath)
	//}

	return nil
}

func (p *CreateAPISubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewAPIScaffolder(p.config, *p.resource, p.force)
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}

func (p *CreateAPISubcommand) PostScaffold() error {
	return nil
}

func (p *CreateAPISubcommand) BindFlags(fs *pflag.FlagSet) {
}
