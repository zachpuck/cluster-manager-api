package cluster_manager_api

import (
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/aws"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/azure"
	"github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api/vmware"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	InvalidProviderErrorMessage           = "valid provider not selected"
	UpgradeNotImplementedErrorMessage     = "upgrades for this cluster/provider not supported yet"
	AdjustNodesNotImplementedErrorMessage = "adjusting nodes for this cluster/provider not supported yet"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	if in.Provider.GetAzure() != nil {
		return s.azure.CreateCluster(in)
	}
	if in.Provider.GetAws() != nil {
		return s.aws.CreateCluster(in)
	}
	if in.Provider.GetVmware() != nil {
		return s.vmware.CreateCluster(in)
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.GetCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.DeleteCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.DeleteCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.DeleteCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	switch in.Provider {
	case pb.Provider_aws:
		if s.aws != nil {
			return s.aws.GetClusterList(in)
		} else {
			return nil, status.Error(codes.Unimplemented, aws.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetClusterList(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetClusterList(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

// Will return upgrade options for a given cluster
func (s *Server) GetUpgradeClusterInformation(ctx context.Context, in *pb.GetUpgradeClusterInformationMsg) (*pb.GetUpgradeClusterInformationReply, error) {
	switch in.Provider {
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.GetClusterUpgrades(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.GetClusterUpgrades(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	case pb.Provider_aws:
		return nil, status.Error(codes.Unimplemented, UpgradeNotImplementedErrorMessage)
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

// Will attempt to upgrade a cluster
func (s *Server) UpgradeCluster(ctx context.Context, in *pb.UpgradeClusterMsg) (*pb.UpgradeClusterReply, error) {
	switch in.Provider {
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.ClusterUpgrade(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.ClusterUpgrade(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	case pb.Provider_aws:
		return nil, status.Error(codes.Unimplemented, UpgradeNotImplementedErrorMessage)
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}

// Will adjust a provision a cluster
func (s *Server) AdjustClusterNodes(ctx context.Context, in *pb.AdjustClusterMsg) (*pb.AdjustClusterReply, error) {

	//TODO: if you pass an empty provider and a valid one program crashes with nil pointer, how better to error for that?

	switch in.Provider {
	case pb.Provider_vmware:
		if s.vmware != nil {
			return s.vmware.AdjustCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, vmware.NotEnabledErrorMessage)
		}
	case pb.Provider_azure:
		if s.azure != nil {
			return s.azure.AdjustCluster(in)
		} else {
			return nil, status.Error(codes.Unimplemented, azure.NotEnabledErrorMessage)
		}

	case pb.Provider_aws:
		return nil, status.Error(codes.Unimplemented, AdjustNodesNotImplementedErrorMessage)
	}
	return nil, status.Error(codes.InvalidArgument, InvalidProviderErrorMessage)
}
