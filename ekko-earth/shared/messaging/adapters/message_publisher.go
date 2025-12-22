package adapters

import (
	"context"

	"github.com/ekko-earth/shared/messaging"
)

type MessagePublisherConfiguration struct {
	Destination string
}

type MessagePublisher interface {
	Publish(message messaging.HasMessageType, context context.Context) error
}
