package adapters

type MessageBus interface {
	Connect() error
	Disconnect() error
}
