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
	"errors"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"csibuilder/pkg/config"
	yamlstore "csibuilder/pkg/config/store"
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins"
)

type CreateCmd struct {
	Command *cobra.Command
	fs      machinery.Filesystem
}

func newCreateCmd() *CreateCmd {
	fs := machinery.Filesystem{FS: afero.NewOsFs()}
	createCmd := &CreateCmd{
		Command: &cobra.Command{
			Use:        "create",
			SuggestFor: []string{"new"},
			Short:      "Scaffold a CSI Driver",
			Long:       `Scaffold a CSI Driver.`,
		},
		fs: fs,
	}

	var options *resourceOptions
	options = bindResourceFlags(createCmd.Command.Flags())
	res := &model.Resource{}
	api := plugins.CreateAPISubcommand{}
	api.BindFlags(createCmd.Command.Flags())

	createCmd.Command.PreRunE = func(cmd *cobra.Command, args []string) error {
		// get conf from config file
		conf, err := createCmd.getInfo()
		if err != nil {
			return err
		}
		if err := api.InjectConfig(conf); err != nil {
			return err
		}
		if options != nil {
			res = options.newResource()
		}
		if err := api.InjectResource(res); err != nil {
			return err
		}
		return api.PreScaffold(fs)
	}
	createCmd.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return api.Scaffold(fs)
	}
	return createCmd
}

func (c *CreateCmd) getInfo() (config.Config, error) {
	conf, err := c.getInfoFromConfigFile()
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	return conf, err
}

// getInfoFromConfigFile obtains the project version and plugin keys from the project config file.
func (c *CreateCmd) getInfoFromConfigFile() (config.Config, error) {
	// Read the project configuration file
	cfg := yamlstore.New(c.fs)
	if err := cfg.Load(); err != nil {
		return nil, err
	}

	return cfg.Config(), nil
}
