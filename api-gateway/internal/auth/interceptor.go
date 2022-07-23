package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/auth/pb"
)

type AuthMiddlewareInterceptor struct {
	*ServiceClient
}

func NewAuthMiddlewareInterceptor(client *ServiceClient) *AuthMiddlewareInterceptor {
	return &AuthMiddlewareInterceptor{client}
}

func (inter *AuthMiddlewareInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		cont, err := inter.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(cont, err)
	}
}

func (inter *AuthMiddlewareInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	if method == "/gophkeeper.grpc.gophkeeper/Login" ||
		method == "/gophkeeper.grpc.gophkeeper/Register" {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.Split(values[0], "Bearer ")
	if len(token) < 2 {
		return ctx, status.Errorf(codes.Unauthenticated, "bad authrization token")
	}

	res, err := inter.Client.Verify(ctx, &pb.VerifyRequest{
		Token: token[1],
	})
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "%v", err)
	}
	if res.Error != "" {
		return ctx, status.Errorf(codes.Internal, "internal error: ", res.Error)
	}

	ctx = context.WithValue(ctx, KeyPrincipalID, res.UserID)
	return ctx, nil
}
