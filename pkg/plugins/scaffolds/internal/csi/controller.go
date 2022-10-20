package csi

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"path/filepath"
)

var _ machinery.Template = &Controller{}

// Controller scaffolds the file that defines the controller service of csi driver
// nolint:maligned
type Controller struct {
	machinery.TemplateMixin
	machinery.MultiGroupMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin

	Force bool
}

func (c *Controller) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join(c.Repo, "pkg/csi/controller.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)
	fmt.Println(c.Path)

	c.TemplateBody = controllerTemplate

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
	}
	return nil
}

const controllerTemplate = `
package csi

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
	}
)

type controllerService struct {
}

func newControllerService() controllerService {
	return controllerService{}
}

func (d *controllerService) CreateVolume(ctx context.Context, request *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if len(request.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume Name cannot be empty")
	}
	if request.VolumeCapabilities == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}

	requiredCap := request.CapacityRange.GetRequiredBytes()

	volCtx := make(map[string]string)
	for k, v := range request.Parameters {
		volCtx[k] = v
	}

	volCtx["subPath"] = request.Name

	volume := csi.Volume{
		VolumeId:      request.Name,
		CapacityBytes: requiredCap,
		VolumeContext: volCtx,
	}

	// TODO modify your createVolume logic here

	return &csi.CreateVolumeResponse{Volume: &volume}, nil
}

func (d *controllerService) DeleteVolume(ctx context.Context, request *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// TODO modify your deleteVolume logic here

	return nil, nil
}

func (d *controllerService) ControllerGetCapabilities(ctx context.Context, request *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	var caps []*csi.ControllerServiceCapability
	for _, cap := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (d *controllerService) ControllerPublishVolume(ctx context.Context, request *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerUnpublishVolume(ctx context.Context, request *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ValidateVolumeCapabilities(ctx context.Context, request *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListVolumes(ctx context.Context, request *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) GetCapacity(ctx context.Context, request *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) CreateSnapshot(ctx context.Context, request *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) DeleteSnapshot(ctx context.Context, request *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListSnapshots(ctx context.Context, request *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerExpandVolume(ctx context.Context, request *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerGetVolume(ctx context.Context, request *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

`
