package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ekko-earth/impact/internal/organisation"

	adapters "github.com/ekko-earth/shared/adapters"
	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"
)

func main() {
	database := gormAdapters.NewGormDatabase(adapters.DatabaseConfiguration{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "ekko_earth",
		Schema:   "impact",
	})

	inboundMessageBus := rabbitmqAdapters.NewRabbitMQMessageBus(rabbitmqAdapters.RabbitMQMessageBusConfiguration{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	})

	organisationDomain := organisation.NewOrganisationDomain(
		database,
		inboundMessageBus,
	)

	organisationDomain.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	defer organisationDomain.Stop()
}
