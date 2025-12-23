package adapters

import (
	"context"
)

type MessagePublisherConfiguration struct{}

type MessagePublisher interface {
	Publish(message any, topic string, ctx context.Context) error
}
