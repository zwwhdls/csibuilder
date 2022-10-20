package csi

import (
	"csibuilder/pkg/machinery"
	"fmt"
	"path/filepath"
)

var _ machinery.Template = &Node{}

// Node scaffolds the file that defines the node server of csi
// nolint:maligned
type Node struct {
	machinery.TemplateMixin
	machinery.MultiGroupMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin

	Force bool
}

func (c *Node) SetTemplateDefaults() error {
	if c.Path == "" {
		c.Path = filepath.Join(c.Repo, "pkg/csi", "node.go")
	}
	c.Path = c.Resource.Replacer().Replace(c.Path)
	fmt.Println(c.Path)

	c.TemplateBody = nodeTemplate

	if c.Force {
		c.IfExistsAction = machinery.OverwriteFile
	} else {
		c.IfExistsAction = machinery.Error
	}
	return nil
}

const nodeTemplate = `
package csi

import (
	"context"
	
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	volumeCaps = []csi.VolumeCapability_AccessMode{
		{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		},
		{
			Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
		},
		{
			Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
		},
	}
)

type nodeService struct {
	nodeID string
}

func newNodeService(nodeID string) nodeService {
	return nodeService{
		nodeID: nodeID,
	}
}

func (n *nodeService) NodeStageVolume(ctx context.Context, request *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (n *nodeService) NodeUnstageVolume(ctx context.Context, request *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (n *nodeService) NodePublishVolume(ctx context.Context, request *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	volumeID := request.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id not provided")
	}

	target := request.GetTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}

	volCap := request.GetVolumeCapability()
	if volCap == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	if !isValidVolumeCapabilities([]*csi.VolumeCapability{volCap}) {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not supported")
	}
	// TODO
	return &csi.NodePublishVolumeResponse{}, nil
}

func (n *nodeService) NodeUnpublishVolume(ctx context.Context, request *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	target := request.GetTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}
	// TODO
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (n *nodeService) NodeGetVolumeStats(ctx context.Context, request *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (n *nodeService) NodeExpandVolume(ctx context.Context, request *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (n *nodeService) NodeGetCapabilities(ctx context.Context, request *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{}, nil
}

func (n *nodeService) NodeGetInfo(ctx context.Context, request *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{NodeId: n.nodeID}, nil
}

func isValidVolumeCapabilities(volCaps []*csi.VolumeCapability) bool {
	hasSupport := func(cap *csi.VolumeCapability) bool {
		for _, c := range volumeCaps {
			if c.GetMode() == cap.AccessMode.GetMode() {
				return true
			}
		}
		return false
	}

	foundAll := true
	for _, c := range volCaps {
		if !hasSupport(c) {
			foundAll = false
		}
	}
	return foundAll
}

`
