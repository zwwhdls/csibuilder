{{ .Boilerplate }}

package csi

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

// GetPluginInfo: returns the name and version of the plugin
func (d *Driver) GetPluginInfo(ctx context.Context, request *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	resp := &csi.GetPluginInfoResponse{
		Name:          DriverName,
		VendorVersion: "v1",
	}
	return resp, nil
}

// GetPluginCapabilities: returns the capabilities of the plugin
func (d *Driver) GetPluginCapabilities(ctx context.Context, request *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	resp := &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
		},
	}
	return resp, nil
}

// Probe: returns the health and readiness of the plugin
func (d *Driver) Probe(ctx context.Context, request *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{}, nil
}
