package main

import (
	"consumers"
	"fmt"
	"log/slog"
	"organisation/internal/adapters"
	"organisation/internal/features/onboard"
	"os"
)

func main() {
	var server consumers.Server

	switch os.Args[1] {
	case "http":
		slog.Info("Creating HTTP server")

		server = adapters.NewHttpServer(adapters.HttpServerConfiguration{
			Address: ":8080",
		})
	case "grpc":
		slog.Info("Creating GRPC server")

		server = adapters.NewGrpcServer(adapters.GrpcServerConfiguration{
			Network: "tcp",
			Port:    50051,
		})
	default:
		panic(fmt.Sprintf("invalid consumer: %s", os.Args[1]))
	}

	onboardFeature := onboard.NewOnboardFeature(server)
	onboardFeature.Start()

	defer onboardFeature.Stop()
}
