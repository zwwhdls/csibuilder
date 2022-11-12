{{ .Boilerplate }}

package main

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
