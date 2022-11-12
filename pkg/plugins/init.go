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
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/spf13/pflag"

	"csibuilder/pkg/config"
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/plugins/scaffolds"
	"csibuilder/pkg/util"
)

type InitSubcommand struct {
	config config.Config
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

func (p *InitSubcommand) InjectConfig(c config.Config) error {
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
	return checkDir()
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

// checkDir will return error if the current directory has files which are not allowed.
// Note that, it is expected that the directory to scaffold the project is cleaned.
// Otherwise, it might face issues to do the scaffold.
func checkDir() error {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Allow directory trees starting with '.'
			if info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			// Allow files starting with '.'
			if strings.HasPrefix(info.Name(), ".") {
				return nil
			}
			// Allow files ending with '.md' extension
			if strings.HasSuffix(info.Name(), ".md") && !info.IsDir() {
				return nil
			}
			// Allow capitalized files except PROJECT
			isCapitalized := true
			for _, l := range info.Name() {
				if !unicode.IsUpper(l) {
					isCapitalized = false
					break
				}
			}
			if isCapitalized && info.Name() != "PROJECT" {
				return nil
			}
			// Allow files in the following list
			allowedFiles := []string{
				"go.mod", // user might run `go mod init` instead of providing the `--flag` at init
				"go.sum", // auto-generated file related to go.mod
			}
			for _, allowedFile := range allowedFiles {
				if info.Name() == allowedFile {
					return nil
				}
			}
			// Do not allow any other file
			return fmt.Errorf(
				"target directory is not empty (only %s, files and directories with the prefix \".\", "+
					"files with the suffix \".md\" or capitalized files name are allowed); "+
					"found existing file %q", strings.Join(allowedFiles, ", "), path)
		})
	if err != nil {
		return err
	}
	return nil
}
