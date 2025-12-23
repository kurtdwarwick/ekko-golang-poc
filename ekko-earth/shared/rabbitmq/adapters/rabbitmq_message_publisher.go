package adapters

import (
	"context"
	"encoding/json"

	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"
	"github.com/google/uuid"

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

func (publisher *RabbitMQMessagePublisher) Publish(message any, topic string, ctx context.Context) error {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	exchange := ""

	if publisher.Configuration.Exchange != nil {
		exchange = *publisher.Configuration.Exchange
	}

	traceId := ctx.Value("traceId")

	if traceId == nil {
		traceId = uuid.New().String()
	}

	channel := publisher.MessageBus.GetChannel(traceId.(string))

	channel.QueueDeclare(
		topic,
		publisher.Configuration.Durable,
		publisher.Configuration.Exclusive,
		publisher.Configuration.AutoDelete,
		publisher.Configuration.NoWait,
		nil,
	)

	err = channel.PublishWithContext(
		ctx,
		exchange,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		panic(err)
	}

	return nil
}
