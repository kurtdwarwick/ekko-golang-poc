package messaging

type MessageHandler[TMessage any] interface {
	Handle(message TMessage) (any, error)
}
