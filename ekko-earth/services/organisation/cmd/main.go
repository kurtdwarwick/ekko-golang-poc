package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ekko-earth/organisation/internal/features/onboard"
	"github.com/ekko-earth/shared/application"
	"github.com/ekko-earth/shared/outbox"

	adapters "github.com/ekko-earth/shared/adapters"
	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	gormOutboxAdapters "github.com/ekko-earth/shared/gorm/adapters/outbox"
	grpcAdapters "github.com/ekko-earth/shared/grpc/adapters"
	httpAdapters "github.com/ekko-earth/shared/http/adapters"
	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"
)

func main() {
	context, cancel := context.WithCancel(context.Background())

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

	outboundMessageBus := rabbitmqAdapters.NewRabbitMQMessageBus(rabbitmqAdapters.RabbitMQMessageBusConfiguration{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	})

	outboundMessagePublisher := rabbitmqAdapters.NewRabbitMQMessagePublisher(
		*outboundMessageBus,
		rabbitmqAdapters.RabbitMQMessagePublisherConfiguration{
			MessagePublisherConfiguration: messagingAdapters.MessagePublisherConfiguration{},
			Durable:                       true,
			Exclusive:                     false,
			AutoDelete:                    false,
			NoWait:                        false,
		},
	)

	unitOfWork := gormAdapters.NewGormUnitOfWork(*database)

	outboxDao := gormOutboxAdapters.NewGormOutboxDAO(*database)
	outboxRepository := outbox.NewOutboxRepository(outboxDao)
	outboxWorker := outbox.NewOutboxWorker(
		outboxRepository,
		unitOfWork,
		outboundMessagePublisher,
		outbox.OutboxWorkerConfiguration{
			PollInterval: time.Second,
			BatchSize:    10000,
			MaxWorkers:   10,
		},
	)

	onboardFeature := onboard.NewOnboardFeature(
		outboundMessagePublisher,
		server,
		database,
		outboxRepository,
		unitOfWork,
	)

	onboardFeature.Start(context)
	outboxWorker.Start(context)

	application.Run(context)

	onboardFeature.Stop(context)
	outboxWorker.Stop(context)

	cancel()
}
