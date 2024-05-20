package interceptor

import (
	"context"
	"hestia/api/pb/hestia/idmanagement"
	"hestia/api/utils/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a gRPC interceptor that checks for a valid token in the metadata
func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Methods to not check
	methods := []string{
		idmanagement.IdentityManagement_Alive_FullMethodName,
		idmanagement.IdentityManagement_Login_FullMethodName,
		idmanagement.IdentityManagement_Logout_FullMethodName,
		idmanagement.IdentityManagement_Register_FullMethodName,
	}

	// Check if the method is in the list of methods to not check
	for _, method := range methods {
		if method == info.FullMethod {
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Missing metadata")
	}

	// Check if the token is in the metadata md["X-AUTH-TOKEN"][0]
	log.Println(md)
	if len(md["x-auth-token"]) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	token := md["x-auth-token"][0]
	if !user.VerifyAuthToken(ctx, token) {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	// continue on
	return handler(ctx, req)

}
