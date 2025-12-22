package adapters

import "context"

type MessageBus interface {
	Connect(context context.Context) error
	Disconnect(context context.Context) error
}
