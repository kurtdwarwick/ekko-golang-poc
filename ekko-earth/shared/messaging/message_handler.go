package messaging

import "context"

type MessageHandler[TMessage any] interface {
	Handle(message TMessage, ctx context.Context) (any, error)
}
