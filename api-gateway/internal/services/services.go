package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/auth"
	authpb "github.com/alexkopcak/gophkeeper/api-gateway/internal/auth/pb"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/command"
	commpb "github.com/alexkopcak/gophkeeper/api-gateway/internal/command/pb"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/query"
	querpb "github.com/alexkopcak/gophkeeper/api-gateway/internal/query/pb"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/services/pb"
)

type APIGatewayService struct {
	pb.UnimplementedAPIGatewayServiceServer
	auth *auth.ServiceClient
	comm *command.ServiceClient
	quer *query.ServiceClient
}

func NewAPIGatewayService(
	authClient *auth.ServiceClient,
	commandClient *command.ServiceClient,
	queryClient *query.ServiceClient,
) pb.APIGatewayServiceServer {
	return &APIGatewayService{
		auth: authClient,
		comm: commandClient,
		quer: queryClient,
	}
}

func (s *APIGatewayService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := s.auth.Client.Register(ctx, &authpb.RegisterRequest{
		UserName: in.UserName,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: res.Token,
		Error: res.Error,
	}, err
}

func (s *APIGatewayService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := s.auth.Client.Login(ctx, &authpb.LoginRequest{
		UserName: in.UserName,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: res.Token,
		Error: res.Error,
	}, err
}

func (s *APIGatewayService) Verify(ctx context.Context, token string) (int64, error) {
	res, err := s.auth.Client.Verify(ctx, &authpb.VerifyRequest{
		Token: token,
	})
	if err != nil {
		return 0, err
	}

	return res.UserID, err
}

func getUserIDfromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(auth.KeyPrincipalID).(int64)
	if !ok {
		return 0, status.Errorf(codes.Unauthenticated, "unknown user")
	}
	return userID, nil
}

func (s *APIGatewayService) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	userID, err := getUserIDfromContext(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.quer.Client.Query(ctx, &querpb.QueryRequest{
		Type:   querpb.MessageType(in.Type),
		UserID: userID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "query error:", err)
	}
	return &pb.QueryResponse{
		Id:    res.Id,
		Data:  res.Data,
		Meta:  res.Meta,
		Error: res.Error,
	}, nil
}

func (s *APIGatewayService) Command(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	userID, err := getUserIDfromContext(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.comm.Client.Command(ctx, &commpb.CommandRequest{
		Operation: commpb.OperationType(in.Operation),
		Type:      commpb.MessageType(in.Type),
		UserID:    userID,
		Data:      in.Data,
		Meta:      in.Meta,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "command error:", err)
	}

	return &pb.CommandResponse{Error: res.Error}, nil
}
