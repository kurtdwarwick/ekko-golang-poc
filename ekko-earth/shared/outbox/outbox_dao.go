package outbox

import (
	"context"

	"github.com/ekko-earth/shared/adapters"
	"github.com/google/uuid"
)

type OutboxDAO interface {
	Create(outboxMessage *OutboxMessage, transaction adapters.Transaction, ctx context.Context) error
	Delete(id uuid.UUID, transaction adapters.Transaction, ctx context.Context) error
	Update(
		where map[string]any,
		update *OutboxMessage,
		limit int,
		transaction adapters.Transaction,
		ctx context.Context,
	) error
	Find(
		where map[string]any,
		limit int,
		transaction adapters.Transaction,
		ctx context.Context,
	) ([]OutboxMessage, error)
}
