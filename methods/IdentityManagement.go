package methods

import (
	"context"

	"hestia/api/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityManagementServer struct {
	pb.UnimplementedIdentityManagementServiceServer
}

func (s *IdentityManagementServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing password")
	}

	token := uuid.New().String()
	return &pb.LoginResponse{Token: token, Name: in.GetEmail(), Email: "a@a.com"}, nil
}

func (s *IdentityManagementServer) Alive(ctx context.Context, in *pb.AliveRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
