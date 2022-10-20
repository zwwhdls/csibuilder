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
		CSIName: opts.CSIName,
	}
}
