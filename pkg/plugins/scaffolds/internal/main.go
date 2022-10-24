package internal

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"os"
	"path/filepath"
)

var _ machinery.Template = &Main{}

// Main scaffolds the main.go
// nolint:maligned
type Main struct {
	machinery.TemplateMixin
	machinery.MultiGroupMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin

	Force bool
}

func (f *Main) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(f.Repo, "main.go")
	}
	f.Path = f.Resource.Replacer().Replace(f.Path)
	fmt.Println(f.Path)

	if f.TemplatePath == "" {
		return fmt.Errorf("can not get template path")
	}

	templateFile := filepath.Join(f.TemplatePath, "main.go.tpl")
	body, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	f.TemplateBody = string(body)

	if f.Force {
		f.IfExistsAction = machinery.OverwriteFile
	} else {
		f.IfExistsAction = machinery.Error
	}
	return nil
}

const mainTemplate = `
package csi

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/klog"

	"{{ .Repo }}/pkg/csi"
)

var (
	endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI Endpoint")
	version  = flag.Bool("version", false, "Print the version and exit.")
	nodeID   = flag.String("nodeid", "", "Node ID")
)

func main() {
	if *version {
		info, err := csi.GetVersionJSON()
		if err != nil {
			klog.Fatalln(err)
		}
		fmt.Println(info)
		os.Exit(0)
	}
	if *nodeID == "" {
		klog.Fatalln("nodeID must be provided")
	}

	drv := csi.NewDriver(*endpoint, *nodeID)
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}
}

`
