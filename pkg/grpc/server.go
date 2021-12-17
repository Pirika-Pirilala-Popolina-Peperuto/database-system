package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	viper.SetDefault("grpc.port", 50051)
}

func NewServerWithLC(lc fx.Lifecycle) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
	)

	reflection.Register(server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			serverAddr := fmt.Sprintf(":%d", viper.GetInt("grpc.port"))
			go func() {
				listener, err := net.Listen("tcp", serverAddr)
				if err != nil {
					log.Fatalf("failed to listen: %s", err)
				}
				if err := server.Serve(listener); err != nil {
					log.Fatalf("failed to serve: %s", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})

	return server
}
