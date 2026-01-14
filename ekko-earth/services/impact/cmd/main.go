package main

import (
	"context"

	"github.com/ekko-earth/impact/internal/organisation"
	"github.com/ekko-earth/shared/application"

	mongoAdapters "github.com/ekko-earth/shared/mongodb/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"
)

func main() {
	context, cancel := context.WithCancel(context.Background())

	database := mongoAdapters.NewMongoDatabase(mongoAdapters.MongoDatabaseConfiguration{
		Host:     "localhost",
		Port:     27017,
		Username: "root",
		Password: "root",
		Database: "ekko_earth",
	})

	// Simply use this to swap out the MongoDB database for a PostgreSQL database

	// database := gormAdapters.NewGormDatabase(adapters.DatabaseConfiguration{
	// 	Host:     "localhost",
	// 	Port:     5432,
	// 	Username: "postgres",
	// 	Password: "postgres",
	// 	Database: "ekko_earth",
	// 	Schema:   "impact",
	// })

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

	organisationDomain.Start(context)

	application.Run(context)

	defer organisationDomain.Stop(context)
	defer cancel()
}
