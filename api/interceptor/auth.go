package interceptor

import (
	"context"
	"hestia/api/pb/idmanagement"
	"hestia/api/utils/user"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
		st := status.New(codes.FailedPrecondition, "missing metadata")

		// Get the service by removing the last part of the full method name
		service := info.FullMethod[:strings.LastIndex(info.FullMethod, "/")]
		ds, err := st.WithDetails(&errdetails.ErrorInfo{
			Reason: "MISSING_METADATA",
			Domain: service,
		})
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}

	if len(md["authorization"]) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	token := md["authorization"][0]
	token_uuid, err := uuid.Parse(token)
	if err != nil {

		return nil, status.Error(codes.InvalidArgument, "Invalid token")
	}
	valid, expired := user.VerifyAuthToken(ctx, token_uuid)
	if expired {
		return nil, status.Error(codes.Unauthenticated, "Token expired")
	} else if !valid {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	// continue on
	return handler(ctx, req)

}
