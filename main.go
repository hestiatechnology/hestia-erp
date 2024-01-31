package main

import (
	"context"
	"log"
	"net"

	pb "hestia/api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHestiaServiceServer
}

func (s *server) GetAddress(ctx context.Context, in *pb.Address) (*pb.Address, error) {
	// Implement your method here.
	// The following is a dummy implementation.
	//log.Printf("Received: %v", in.GetYourField())
	return &pb.Address{City: "Gay", Street: "aasd", State: "sdfsdf", Country: "safdf"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterHestiaServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
