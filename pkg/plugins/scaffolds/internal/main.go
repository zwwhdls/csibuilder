package internal

import (
	"csibuilder/pkg/machinery"
	"fmt"
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

func (c *Main) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join(c.Repo, "main.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)
	fmt.Println(c.Path)

	c.TemplateBody = mainTemplate

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
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
