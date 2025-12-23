package outbox

import (
	"time"

	"github.com/google/uuid"
)

type OutboxMessage struct {
	Id            uuid.UUID
	MessageType   string
	Message       any
	CreatedAt     time.Time
	LockedAt      *time.Time
	LockReference *string
}

func (outboxMessage *OutboxMessage) GetMessageType() string {
	return outboxMessage.MessageType
}
