package adapters

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/ekko-earth/shared/messaging"

	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMessagePublisher struct {
	Configuration RabbitMQMessagePublisherConfiguration
	MessageBus    RabbitMQMessageBus
}

type RabbitMQMessagePublisherConfiguration struct {
	messagingAdapters.MessagePublisherConfiguration

	Exchange *string

	Durable    bool
	Exclusive  bool
	AutoDelete bool
	NoWait     bool
}

func NewRabbitMQMessagePublisher(
	messageBus RabbitMQMessageBus,
	configuration RabbitMQMessagePublisherConfiguration,
) *RabbitMQMessagePublisher {
	return &RabbitMQMessagePublisher{
		MessageBus:    messageBus,
		Configuration: configuration,
	}
}

func (publisher *RabbitMQMessagePublisher) Publish(message messaging.HasMessageType, context context.Context) error {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	exchange := ""

	if publisher.Configuration.Exchange != nil {
		exchange = *publisher.Configuration.Exchange
	}

	slog.Info("Declaring queue", "destination", message.GetMessageType())
	slog.Info("Message", "message", message)

	publisher.MessageBus.Channel.QueueDeclare(
		message.GetMessageType(),
		publisher.Configuration.Durable,
		publisher.Configuration.Exclusive,
		publisher.Configuration.AutoDelete,
		publisher.Configuration.NoWait,
		nil,
	)

	err = publisher.MessageBus.Channel.PublishWithContext(
		context,
		exchange,
		message.GetMessageType(),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
