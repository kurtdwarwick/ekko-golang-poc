package rabbitmq

import (
	"github.com/ekko-earth/impact/internal/organisation/core/events"
	"github.com/ekko-earth/shared/messaging"
)

type OrganisationOnboardedEventMessageTranslator struct {
	MessageTranslator messaging.MessageTranslator[events.OrganisationOnboardedEvent, events.OrganisationOnboardedEvent]
}

func (translator *OrganisationOnboardedEventMessageTranslator) Translate(
	message events.OrganisationOnboardedEvent,
) (events.OrganisationOnboardedEvent, error) {
	return message, nil
}
