package main

import (
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	repo = ""
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "create",
		SuggestFor: []string{"new"},
		Short:      "Scaffold a Kubernetes API",
		Long:       `Scaffold a Kubernetes API.`,
	}

	var options *resourceOptions
	options = bindResourceFlags(cmd.Flags())
	var res *model.Resource
	api := plugins.CreateAPISubcommand{}
	api.BindFlags(cmd.Flags())

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if options != nil {
			res = options.newResource()
		}
		return nil
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		fs := machinery.Filesystem{FS: afero.NewOsFs()}
		if err := api.InjectResource(res); err != nil {
			return err
		}
		return api.Scaffold(fs)
	}
	return cmd
}
