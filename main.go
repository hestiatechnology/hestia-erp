package main

import (
	"net"
	"os"
	"strconv"
	"strings"

	"hestia/api/interceptor"
	"hestia/api/methods"
	"hestia/api/pb/company"
	"hestia/api/pb/idmanagement"
	"hestia/api/pb/textile"
	"hestia/api/utils/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	PORT := 9000
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		logger.ErrorLogger.Fatalf("failed to listen: %v", err)
	}
	logger.InfoLogger.Println("Server listening on port", PORT)
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor))

	if strings.ToLower(os.Getenv("ENV")) == "development" {
		logger.WarningLogger.Println("Running in development mode")
		logger.WarningLogger.Println("Registering reflection service")
		reflection.Register(s)
	}

	// Service registration
	idmanagement.RegisterIdentityManagementServer(s, &methods.IdentityManagementServer{})
	textile.RegisterTextileServer(s, &methods.TextileServer{})
	company.RegisterCompanyManagementServer(s, &methods.CompanyManagementServer{})
	if err := s.Serve(lis); err != nil {
		logger.ErrorLogger.Fatalf("failed to serve: %v", err)
	}
}
