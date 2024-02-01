package methods

import (
	"context"

	pb "hestia/api/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IdentityManagementServer struct {
	pb.UnimplementedIdentityManagementServiceServer
}

func (s *IdentityManagementServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing email")
	}
	return &pb.LoginResponse{Token: in.Email, Name: in.GetEmail(), Email: "a@a.com"}, nil
}
