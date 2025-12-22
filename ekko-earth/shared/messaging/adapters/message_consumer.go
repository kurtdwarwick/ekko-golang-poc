package adapters

type MessageConsumer[TMessage any] interface {
	Consume(message TMessage) error
}
