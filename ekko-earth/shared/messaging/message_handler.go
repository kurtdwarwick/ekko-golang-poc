package messaging

import "context"

type MessageHandler[TMessage any] interface {
	Handle(message TMessage, context context.Context) (any, error)
}
