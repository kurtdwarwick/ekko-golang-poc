package adapters

import "context"

type MessageConsumer[TMessage any] interface {
	Consume(message TMessage, ctx context.Context) error
}
