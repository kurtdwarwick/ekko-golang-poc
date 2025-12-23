package adapters

import (
	"context"
	"encoding/json"

	"github.com/ekko-earth/shared/messaging"
)

type RabbitMQMessageConsumer[TIncomingMessage any, TMessage any] struct {
	messageHandler messaging.MessageHandler[TMessage]
}

type RabbitMQMessageConsumerConfiguration struct {
	Queue           string
	AutoAcknowledge bool
}

func NewRabbitMQMessageConsumer[TIncomingMessage any, TMessage any](
	messageBus RabbitMQMessageBus,
	messageTranslator messaging.MessageTranslator[TIncomingMessage, TMessage],
	messageHandler messaging.MessageHandler[TMessage],
	configuration RabbitMQMessageConsumerConfiguration,
) *RabbitMQMessageConsumer[TIncomingMessage, TMessage] {
	channel := messageBus.GetChannel(configuration.Queue)

	channel.QueueDeclare(
		configuration.Queue,
		true,
		false,
		false,
		false,
		nil,
	)

	deliveries, err := channel.Consume(
		configuration.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	consumer := &RabbitMQMessageConsumer[TIncomingMessage, TMessage]{
		messageHandler: messageHandler,
	}

	go func() {
		for d := range deliveries {
			var incomingMessage TIncomingMessage

			err := json.Unmarshal(d.Body, &incomingMessage)

			if err != nil {
				panic(err)
			}

			translatedMessage, err := messageTranslator.Translate(incomingMessage)

			if err != nil {
				panic(err)
			}

			context, cancel := context.WithCancel(context.Background())

			messageHandler.Handle(translatedMessage, context)

			cancel()
		}
	}()

	return consumer
}
