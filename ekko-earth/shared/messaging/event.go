package messaging

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Message

	Id         uuid.UUID
	OccurredAt time.Time
}
