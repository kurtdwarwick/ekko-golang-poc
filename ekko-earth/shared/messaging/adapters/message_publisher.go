package adapters

import (
	"context"
)

type MessagePublisherConfiguration struct{}

type MessagePublisher interface {
	Publish(message any, topic string, headers map[string]any, ctx context.Context) error
}
