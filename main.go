package main

import (
	"hestia/api/utils/log"
	"net"
	"os"
	"strconv"
	"strings"

	"hestia/api/interceptor"
	"hestia/api/methods"
	"hestia/api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	PORT := 9000
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.ErrorLogger.Fatalf("failed to listen: %v", err)
	}
	log.InfoLogger.Println("Server listening on port", PORT)
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor))

	if strings.ToLower(os.Getenv("ENV")) == "development" {
		log.WarningLogger.Println("Running in development mode")
		log.WarningLogger.Println("Registering reflection service")
		reflection.Register(s)
	}

	// Service registration
	pb.RegisterIdentityManagementServer(s, &methods.IdentityManagementServer{})
	pb.RegisterTextileServer(s, &methods.TextileServer{})
	pb.RegisterCompanyManagementServer(s, &methods.CompanyManagementServer{})
	if err := s.Serve(lis); err != nil {
		log.ErrorLogger.Fatalf("failed to serve: %v", err)
	}
}
