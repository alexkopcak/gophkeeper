package command

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/command/pb"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/config"
)

type ServiceClient struct {
	Client pb.CommandServiceClient
}

func InitCommandServiceClient(ctx context.Context, cfg *config.Config) pb.CommandServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.CommandServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect: ", err)
	}
	defer cc.Close()

	return pb.NewCommandServiceClient(cc)
}

func NewCommandServiceClient(ctx context.Context, cfg *config.Config) *ServiceClient {
	return &ServiceClient{
		Client: InitCommandServiceClient(ctx, cfg),
	}
}
