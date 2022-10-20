package csi

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"path/filepath"
)

var _ machinery.Template = &Driver{}

// Driver scaffolds the file that defines csi driver
// nolint:maligned
type Driver struct {
	machinery.TemplateMixin
	machinery.MultiGroupMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin

	Force bool
}

func (c *Driver) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join(c.Repo, "pkg/csi", "driver.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)
	fmt.Println(c.Path)

	c.TemplateBody = driverTemplate

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
	}
	return nil
}

const driverTemplate = `
package csi

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"k8s.io/klog"
)

const (
	// DriverName to be registered
	DriverName = "{{ .Resource.CSIName }}"
)

type Driver struct {
	controllerService
	nodeService

	srv      *grpc.Server
	endpoint string
}

// NewDriver creates a new driver
func NewDriver(endpoint string, nodeID string) *Driver {
	klog.Infof("Driver: %v version %v commit %v date %v", DriverName, driverVersion, gitCommit, buildDate)

	return &Driver{
		endpoint:          endpoint,
		controllerService: newControllerService(),
		nodeService:       newNodeService(nodeID),
	}
}

func (d *Driver) Run() error {
	scheme, addr, err := ParseEndpoint(d.endpoint)
	if err != nil {
		return err
	}

	listener, err := net.Listen(scheme, addr)
	if err != nil {
		return err
	}

	logErr := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			klog.Errorf("GRPC error: %v", err)
		}
		return resp, err
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logErr),
	}
	d.srv = grpc.NewServer(opts...)

	csi.RegisterIdentityServer(d.srv, d)
	csi.RegisterControllerServer(d.srv, d)
	csi.RegisterNodeServer(d.srv, d)

	klog.Infof("Listening for connection on address: %#v", listener.Addr())
	return d.srv.Serve(listener)
}

func ParseEndpoint(endpoint string) (string, string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", fmt.Errorf("could not parse endpoint: %v", err)
	}

	addr := path.Join(u.Host, filepath.FromSlash(u.Path))

	scheme := strings.ToLower(u.Scheme)
	switch scheme {
	case "tcp":
	case "unix":
		addr = path.Join("/", addr)
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			return "", "", fmt.Errorf("could not remove unix domain socket %q: %v", addr, err)
		}
	default:
		return "", "", fmt.Errorf("unsupported protocol: %s", scheme)
	}

	return scheme, addr, nil
}

`
