package adapters

import (
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

func (server *GrpcServer) Start() {
	slog.Info("Starting GRPC consumer")

	server.Listener.Accept()
}

func (server *GrpcServer) Stop() {
	slog.Info("Stopping GRPC consumer")

	server.Listener.Close()
}
