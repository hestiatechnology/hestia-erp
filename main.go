package main

import (
	"log"
	"net"
	"os"
	"strings"

	"hestia/api/interceptor"
	"hestia/api/methods"
	"hestia/api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor))

	if strings.ToLower(os.Getenv("ENV")) == "development" {
		log.Println("Running in development mode")
		log.Println("Registering reflection service")
		reflection.Register(s)
	}

	// Service registration
	pb.RegisterIdentityManagementServiceServer(s, &methods.IdentityManagementServer{})
	pb.RegisterTextileServiceServer(s, &methods.TextileServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
