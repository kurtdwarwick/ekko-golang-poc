package adapters

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDatabase struct {
	Configuration MongoDatabaseConfiguration
	Client        *mongo.Client
}

type MongoDatabaseConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewMongoDatabase(configuration MongoDatabaseConfiguration) *MongoDatabase {
	address := net.JoinHostPort(configuration.Host, strconv.Itoa(configuration.Port))
	host := fmt.Sprintf(
		"mongodb://%s:%s@%s/%s/?authSource=admin",
		configuration.Username,
		configuration.Password,
		address,
		configuration.Database,
	)

	slog.Info("Connecting to MongoDB", "host", host)

	client, err := mongo.Connect(options.Client().ApplyURI(host))

	if err != nil {
		panic(err)
	}

	return &MongoDatabase{Configuration: configuration, Client: client}
}

func (database *MongoDatabase) Connect() error {
	return nil
}

func (database *MongoDatabase) Disconnect() error {
	err := database.Client.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	return nil
}
