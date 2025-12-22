package events

import (
	"github.com/ekko-earth/shared/messaging"
	"github.com/google/uuid"
)

type OrganisationOnboardedEvent struct {
	messaging.Event

	OrganisationId uuid.UUID
}
