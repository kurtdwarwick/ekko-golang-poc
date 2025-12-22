package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ekko-earth/organisation/internal/features/onboard"
	"github.com/ekko-earth/shared/application"

	adapters "github.com/ekko-earth/shared/adapters"
	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	grpcAdapters "github.com/ekko-earth/shared/grpc/adapters"
	httpAdapters "github.com/ekko-earth/shared/http/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"
)

func main() {
	var server adapters.Server

	switch os.Args[1] {
	case "http":
		slog.Info("Creating HTTP server")

		server = httpAdapters.NewHttpServer(httpAdapters.HttpServerConfiguration{
			Address: ":8080",
		})
	case "grpc":
		slog.Info("Creating GRPC server")

		server = grpcAdapters.NewGrpcServer(grpcAdapters.GrpcServerConfiguration{
			Network: "tcp",
			Port:    50051,
		})
	default:
		panic(fmt.Sprintf("invalid consumer: %s", os.Args[1]))
	}

	database := gormAdapters.NewGormDatabase(adapters.DatabaseConfiguration{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "ekko_earth",
		Schema:   "organisation",
	})

	inboundMessageBus := rabbitmqAdapters.NewRabbitMQMessageBus(rabbitmqAdapters.RabbitMQMessageBusConfiguration{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	})

	outboundMessageBus := rabbitmqAdapters.NewRabbitMQMessageBus(rabbitmqAdapters.RabbitMQMessageBusConfiguration{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	})

	onboardFeature := onboard.NewOnboardFeature(
		inboundMessageBus,
		outboundMessageBus,
		server,
		database,
	)

	onboardFeature.Start()

	application.Run()

	defer onboardFeature.Stop()
}
