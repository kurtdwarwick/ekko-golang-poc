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
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

func NewRabbitMQMessageBus(configuration RabbitMQMessageBusConfiguration) *RabbitMQMessageBus {
	address := net.JoinHostPort(configuration.Host, strconv.Itoa(configuration.Port))
	host := fmt.Sprintf("amqp://%s:%s@%s", configuration.Username, configuration.Password, address)

	slog.Info("Connecting to RabbitMQ", "host", host)

	connection, err := amqp.Dial(host)

	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()

	if err != nil {
		panic(err)
	}

	return &RabbitMQMessageBus{
		Connection: connection,
		Channel:    channel,
	}
}

func (bus *RabbitMQMessageBus) Connect(ctx context.Context) error {
	return nil
}

func (bus *RabbitMQMessageBus) Disconnect(ctx context.Context) error {
	bus.Connection.Close()
	bus.Channel.Close()

	return nil
}
