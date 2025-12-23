package adapters

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMessageBusConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
}

type RabbitMQMessageBus struct {
	Channels   map[string]*amqp.Channel
	Connection *amqp.Connection
}

func NewRabbitMQMessageBus(configuration RabbitMQMessageBusConfiguration) *RabbitMQMessageBus {
	address := net.JoinHostPort(configuration.Host, strconv.Itoa(configuration.Port))
	host := fmt.Sprintf("amqp://%s:%s@%s", configuration.Username, configuration.Password, address)

	slog.Info("Connecting to RabbitMQ")

	connection, err := amqp.Dial(host)

	if err != nil {
		panic(err)
	}

	return &RabbitMQMessageBus{
		Connection: connection,
		Channels:   make(map[string]*amqp.Channel),
	}
}

func (bus *RabbitMQMessageBus) GetChannel(key string) *amqp.Channel {
	if _, ok := bus.Channels[key]; !ok {
		channel, err := bus.Connection.Channel()

		if err != nil {
			panic(err)
		}

		bus.Channels[key] = channel
	}

	return bus.Channels[key]
}

func (bus *RabbitMQMessageBus) Connect(ctx context.Context) error {
	return nil
}

func (bus *RabbitMQMessageBus) Disconnect(ctx context.Context) error {
	bus.Connection.Close()

	for _, channel := range bus.Channels {
		channel.Close()
	}

	return nil
}
