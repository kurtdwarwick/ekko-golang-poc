package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/ekko-earth/organisation/internal/features/onboard"
	"github.com/ekko-earth/organisation/internal/features/query"
	"github.com/ekko-earth/shared/application"

	adapters "github.com/ekko-earth/shared/adapters"
	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	http "github.com/ekko-earth/shared/http"
	httpAdapters "github.com/ekko-earth/shared/http/adapters"
	observability "github.com/ekko-earth/shared/observability"
	outbox "github.com/ekko-earth/shared/outbox"
	outboxAdapters "github.com/ekko-earth/shared/outbox/adapters/gorm"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"
)

func main() {
	context, cancel := context.WithCancel(context.Background())
	shutdown, err := observability.NewInstrumentation(context, "github.com/ekko-earth/organisation")

	var server adapters.Server

	if err != nil {
		slog.Error("Failed to configure HTTP instrumenter", "error", err)
		panic(err)
	}

	defer shutdown(context)

	slog.Info("Creating HTTP server")

	server = httpAdapters.NewHttpServer(httpAdapters.HttpServerConfiguration{
		Address: ":8080",
	})

	// Simply use this to swap out the HTTP server for a GRPC server

	// slog.Info("Creating GRPC server")

	// server = grpcAdapters.NewGrpcServer(grpcAdapters.GrpcServerConfiguration{
	// 	Network: "tcp",
	// 	Port:    50051,
	// })

	http.Instrument(*server.(*httpAdapters.HttpServer))

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
			Durable:    true,
			Exclusive:  false,
			AutoDelete: false,
			NoWait:     false,
		},
	)

	unitOfWork := gormAdapters.NewGormUnitOfWork(*database)

	outboxDao := outboxAdapters.NewGormOutboxDAO(*database)
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

	queryFeature := query.NewQueryFeature(
		server,
		database,
	)

	server.Start(context)
	database.Connect(context)

	onboardFeature.Start(context)
	queryFeature.Start(context)
	outboxWorker.Start(context)

	application.Run(context)

	onboardFeature.Stop(context)
	queryFeature.Stop(context)
	outboxWorker.Stop(context)

	cancel()
}
