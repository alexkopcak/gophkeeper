package query

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/config"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/query/pb"
)

type ServiceClient struct {
	Client pb.QueryServiceClient
}

func InitQueryServiceClient(ctx context.Context, cfg *config.Config) pb.QueryServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.QueryServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer cc.Close()

	return pb.NewQueryServiceClient(cc)
}

func NewQueryServiceClient(ctx context.Context, cfg *config.Config) *ServiceClient {
	return &ServiceClient{
		Client: InitQueryServiceClient(ctx, cfg),
	}
}
