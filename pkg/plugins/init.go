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
	"path/filepath"

	"github.com/spf13/pflag"

	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins/scaffolds"
	"csibuilder/pkg/util"
)

type InitSubcommand struct {
	config model.Config
	// For help text.
	commandName string

	// boilerplate options
	license string
	owner   string

	// go config options
	repo string

	// flags
	fetchDeps          bool
	skipGoVersionCheck bool
}

func (p *InitSubcommand) BindFlags(fs *pflag.FlagSet) {
	// dependency args
	fs.BoolVar(&p.fetchDeps, "fetch-deps", true, "ensure dependencies are downloaded")

	// boilerplate args
	fs.StringVar(&p.license, "license", "apache2",
		"license to use to boilerplate, may be one of 'apache2', 'none'")
	fs.StringVar(&p.owner, "owner", "", "owner to add to the copyright")

	// project args
	fs.StringVar(&p.repo, "repo", "", "name to use for go module (e.g., github.com/user/repo), "+
		"defaults to the go package of the current working directory.")
}

func (p *InitSubcommand) InjectConfig(c model.Config) error {
	p.config = c

	// Try to guess repository if flag is not set.
	if p.repo == "" {
		repoPath, err := util.FindCurrentRepo()
		if err != nil {
			return fmt.Errorf("error finding current repository: %v", err)
		}
		p.repo = repoPath
	}

	return p.config.SetRepository(p.repo)
}

func (p *InitSubcommand) PreScaffold(machinery.Filesystem) error {
	// Check if the current directory has not files or directories which does not allow to init the project
	// inject template path
	curPath, err := filepath.Abs("")
	if err != nil {
		return fmt.Errorf("can not get abs path: %s", err)
	}
	return p.config.SetTemplatePath(filepath.Join(curPath, "pkg/templates"))
}

func (p *InitSubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewInitScaffolder(p.config, p.license, p.owner)
	scaffolder.InjectFS(fs)
	err := scaffolder.Scaffold()
	if err != nil {
		return err
	}

	if !p.fetchDeps {
		fmt.Println("Skipping fetching dependencies.")
		return nil
	}

	return nil
}

func (p *InitSubcommand) PostScaffold() error {
	err := util.RunCmd("Update dependencies", "go", "mod", "tidy")
	if err != nil {
		return err
	}

	fmt.Printf("Next: define a csi with:\n$ %s create api\n", p.commandName)
	return nil
}
