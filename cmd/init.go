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
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "init",
		SuggestFor: []string{"new"},
		Short:      "Scaffold a CSI Driver",
		Long:       `Scaffold a CSI Driver.`,
	}

	var (
		conf model.Config
	)
	api := plugins.InitSubcommand{}
	api.BindFlags(cmd.Flags())
	fs := machinery.Filesystem{FS: afero.NewOsFs()}

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if err := api.InjectConfig(conf); err != nil {
			return err
		}
		return api.PreScaffold(fs)
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return api.Scaffold(fs)
	}
	return cmd
}
