package services

import (
	"context"

	commandpb "github.com/alexkopcak/gophkeeper/api-gateway/pkg/command/pb"
)

type CommandServer struct {
	commandpb.UnimplementedCommandServiceServer
}

var _ commandpb.CommandServiceServer = (*CommandServer)(nil)

func (s *CommandServer) Command(ctx context.Context, in *commandpb.CommandRequest) (*commandpb.CommandResponse, error) {

	return nil, nil
}
