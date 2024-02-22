package methods

import (
	"context"
	"hestia/api/pb"
)

type CompanyManagementServer struct {
	pb.UnimplementedCompanyManagementServer
}

func (s *CompanyManagementServer) CreateCompany(ctx context.Context, in *pb.CreateCompanyRequest) (*pb.IdResponse, error) {
	return &pb.IdResponse{}, nil
}
