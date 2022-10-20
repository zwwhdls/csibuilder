package scaffolds

import (
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins/scaffolds/internal"
	"csibuilder/pkg/plugins/scaffolds/internal/csi"
	"fmt"
)

// ApiScaffolder contains configuration for generating scaffolding for Go type
// representing the API and controller that implements the behavior for the API.
type ApiScaffolder struct {
	// fs is the filesystem that will be used by the scaffolder
	fs machinery.Filesystem

	// force indicates whether to scaffold controller files even if it exists or not
	force    bool
	resource model.Resource
	config   model.Config
}

// NewAPIScaffolder returns a new Scaffolder for API/controller creation operations
func NewAPIScaffolder(conf model.Config, res model.Resource, force bool) *ApiScaffolder {
	return &ApiScaffolder{
		force:    force,
		resource: res,
		config:   conf,
	}
}

func (s *ApiScaffolder) InjectFS(filesystem machinery.Filesystem) {
	s.fs = filesystem
}

func (s *ApiScaffolder) Scaffold() error {
	fmt.Println("Writing scaffold for you to edit...")

	scaffold := machinery.NewScaffold(s.fs,
		machinery.WithResource(&s.resource),
		machinery.WithConfig(&s.config),
	)
	if err := scaffold.Execute(
		&csi.Controller{Force: s.force},
		&csi.Driver{Force: s.force},
		&csi.Identity{Force: s.force},
		&csi.Node{Force: s.force},
		&csi.Version{Force: s.force},
		&internal.Main{Force: s.force},
		&internal.GoMod{},
	); err != nil {
		return fmt.Errorf("error scaffolding APIs: %v", err)
	}

	return nil
}
