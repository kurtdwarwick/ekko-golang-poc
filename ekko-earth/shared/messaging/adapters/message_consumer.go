package adapters

import "context"

type MessageConsumer[TMessage any] interface {
	Consume(message TMessage, context context.Context) error
}
