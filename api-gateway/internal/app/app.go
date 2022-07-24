package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/alexkopcak/gophkeeper/api-gateway/internal/auth"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/command"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/config"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/query"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/services"
	"github.com/alexkopcak/gophkeeper/api-gateway/internal/services/pb"
)

type App struct {
	cfg           *config.Config
	authClient    *auth.ServiceClient
	interceptor   *auth.AuthMiddlewareInterceptor
	queryClient   *query.ServiceClient
	commandClient *command.ServiceClient
	serverGRPC    *grpc.Server
}

const (
	certFile = "localhost.crt"
	keyFile  = "localhost.key"
)

func NewApp() *App {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("failed at config", err)
	}

	return &App{
		cfg: cfg,
	}
}

func (app *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	app.authClient = auth.NewAuthServiceClient(ctx, app.cfg)
	app.interceptor = auth.NewAuthMiddlewareInterceptor(app.authClient)
	app.queryClient = query.NewQueryServiceClient(ctx, app.cfg)
	app.commandClient = command.NewCommandServiceClient(ctx, app.cfg)

	idleConnClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigint
		close(idleConnClosed)
	}()

	go func() {
		<-idleConnClosed
		if app.serverGRPC != nil {
			app.serverGRPC.GracefulStop()
		}
	}()

	return app.startGRPC()
}

func (app *App) startGRPC() error {
	listen, err := net.Listen("tcp", app.cfg.Port)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(app.interceptor.Unary()),
	}
	if app.cfg.TLS {
		creds, er := credentials.NewServerTLSFromFile(certFile, keyFile)
		if er != nil {
			return er
		}
		opts = append(opts, grpc.Creds(creds))
	}

	app.serverGRPC = grpc.NewServer(opts...)

	pb.RegisterAPIGatewayServiceServer(
		app.serverGRPC,
		services.NewAPIGatewayService(
			app.authClient,
			app.commandClient,
			app.queryClient,
		))

	log.Printf("API Gateway service start on %v", app.cfg.Port)

	err = app.serverGRPC.Serve(listen)
	if err == nil {
		log.Println("API Gateway service graceful shutdown")
	}
	return err
}
