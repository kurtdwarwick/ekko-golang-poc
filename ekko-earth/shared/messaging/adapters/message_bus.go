package adapters

import "context"

type MessageBus interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
