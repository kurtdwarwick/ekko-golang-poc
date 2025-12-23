package outbox

import (
	"context"
	"time"

	"github.com/ekko-earth/shared/adapters"
	"github.com/google/uuid"
)

type OutboxRepository struct {
	outboxDao OutboxDAO
}

func NewOutboxRepository(outboxDao OutboxDAO) *OutboxRepository {
	return &OutboxRepository{outboxDao: outboxDao}
}

func (repository *OutboxRepository) ScheduleMessage(
	outboxMessage *OutboxMessage,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	return repository.outboxDao.Create(outboxMessage, transaction, ctx)
}

func (repository *OutboxRepository) LockUnsentMessages(
	reference string,
	batchSize int,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	lockedAt := time.Now()

	return repository.outboxDao.Update(
		map[string]any{"lock_reference": nil},
		&OutboxMessage{LockedAt: &lockedAt, LockReference: &reference},
		batchSize,
		transaction,
		ctx,
	)
}

func (repository *OutboxRepository) GetUnsentMessages(
	reference string,
	batchSize int,
	transaction adapters.Transaction,
	ctx context.Context,
) ([]OutboxMessage, error) {
	return repository.outboxDao.Find(map[string]any{"lock_reference": reference}, batchSize, transaction, ctx)
}

func (repository *OutboxRepository) RemoveMessage(
	id uuid.UUID,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	return repository.outboxDao.Delete(id, transaction, ctx)
}
