package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/auth/pb"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/config"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

type key uint64

const (
	KeyPrincipalID key = iota
)

func NewAuthServiceClient(ctx context.Context, cfg *config.Config) *ServiceClient {
	return &ServiceClient{
		Client: InitAuthServiceClient(ctx, cfg),
	}
}

func InitAuthServiceClient(ctx context.Context, cfg *config.Config) pb.AuthServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.AuthServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	defer cc.Close()

	return pb.NewAuthServiceClient(cc)
}
