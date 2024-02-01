package main

import (
	"log"
	"net"
	"os"

	idm "hestia/api/methods"
	pb "hestia/api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	if os.Getenv("HESTIA_ENV") == "development" {
		log.Println("Running in development mode")
		log.Println("Registering reflection service")
		reflection.Register(s)
	}

	// Service registration
	pb.RegisterIdentityManagementServiceServer(s, &idm.IdentityManagementServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
