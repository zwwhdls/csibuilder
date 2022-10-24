package plugins

import (
	"csibuilder/pkg/machinery"
	"csibuilder/pkg/model"
	"csibuilder/pkg/plugins/scaffolds"
	"fmt"
	"github.com/spf13/pflag"
	"path/filepath"
)

const (
	// defaultCRDVersion is the default CRD API version to scaffold.
	defaultCRDVersion = "v1"
)

// DefaultMainPath is default file path of main.go
const DefaultMainPath = "main.go"

type CreateAPISubcommand struct {
	config   *model.Config
	resource *model.Resource

	// Check if we have to scaffold resource and/or controller
	resourceFlag   *pflag.Flag
	controllerFlag *pflag.Flag

	// force indicates that the resource should be created even if it already exists
	force bool

	// runMake indicates whether to run make or not after scaffolding APIs
	runMake bool
}

func (p *CreateAPISubcommand) InjectResource(res *model.Resource) error {
	p.resource = res

	return nil
}

func (p *CreateAPISubcommand) PreScaffold(machinery.Filesystem) error {
	// check if main.go is present in the root directory
	//if _, err := os.Stat(DefaultMainPath); os.IsNotExist(err) {
	//	return fmt.Errorf("%s file should present in the root directory", DefaultMainPath)
	//}

	// inject template path
	curPath, err := filepath.Abs("")
	if err != nil {
		return fmt.Errorf("can not get abs path: %s", err)
	}
	return p.config.SetTemplatePath(filepath.Join(curPath, "pkg/templates"))
}

func (p *CreateAPISubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewAPIScaffolder(*p.config, *p.resource, p.force)
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}

func (p *CreateAPISubcommand) PostScaffold() error {
	return nil
}

func (p *CreateAPISubcommand) BindFlags(fs *pflag.FlagSet) {
	if p.config == nil {
		p.config = &model.Config{}
	}
	fs.StringVar(&p.config.Repo, "repo", "", "name to use for go module (e.g., github.com/user/repo), "+
		"defaults to the go package of the current working directory.")
}
