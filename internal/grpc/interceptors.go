package grpc

import (
	"context"
	"github.com/babenkoivan/todo/internal/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var permittedMethods = map[string]bool{
	"/UserService/Login":  true,
	"/UserService/Create": true,
}

func authenticate(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if permittedMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	user, err := users.Authenticate(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid")
	}

	ctx = setUser(ctx, user)
	return handler(ctx, req)
}
