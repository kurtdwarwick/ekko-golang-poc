package adapters

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

type GrpcServerConfiguration struct {
	Address string
	Network string

	Port int
}

func NewGrpcServer(configuration GrpcServerConfiguration) *GrpcServer {
	listener, err := net.Listen(configuration.Network, configuration.Address)

	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer()

	return &GrpcServer{Server: server, Listener: listener}
}

func (server *GrpcServer) Start(context context.Context) error {
	slog.Info("Starting GRPC consumer")

	go func() {
		if _, err := server.Listener.Accept(); err != nil {
			slog.Error("Failed to accept", "error", err)
		}
	}()

	return nil
}

func (server *GrpcServer) Stop(context context.Context) error {
	slog.Info("Stopping GRPC consumer")

	server.Listener.Close()

	return nil
}
