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
	fs := machinery.Filesystem{FS: afero.NewOsFs()}

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if options != nil {
			res = options.newResource()
		}
		if err := api.InjectResource(res); err != nil {
			return err
		}
		return api.PreScaffold(fs)
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return api.Scaffold(fs)
	}
	return cmd
}
