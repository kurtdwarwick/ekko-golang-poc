package handlers

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/commands"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/repositories"

	organisationEvents "github.com/ekko-earth/organisation/internal/features/onboard/core/events"
	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"

	"github.com/ekko-earth/shared/messaging"
)

type OnboardOrganisationCommandHandler struct {
	repository       *repositories.OrganisationRepository
	messagePublisher messagingAdapters.MessagePublisher
}

func NewOnboardOrganisationCommandHandler(
	repository *repositories.OrganisationRepository,
	messagePublisher messagingAdapters.MessagePublisher,
) *OnboardOrganisationCommandHandler {
	return &OnboardOrganisationCommandHandler{
		repository:       repository,
		messagePublisher: messagePublisher,
	}
}

func (handler *OnboardOrganisationCommandHandler) Handle(
	command commands.OnboardOrganisationCommand,
	context context.Context,
) (any, error) {
	organisationId, err := handler.repository.OnboardOrganisation(
		entities.Organisation{
			LegalName:   command.LegalName,
			TradingName: command.TradingName,
			Website:     command.Website,
		}, context)

	if err != nil {
		slog.Error("Failed to onboard organisation", "error", err)
		return nil, err
	}

	event := organisationEvents.OrganisationOnboardedEvent{
		Event: messaging.Event{
			Message: messaging.Message{
				ConversationId: command.Message.ConversationId,
			},
			Id:         uuid.New(),
			OccurredAt: time.Now(),
		},
		OrganisationId: *organisationId,
		LegalName:      command.LegalName,
		TradingName:    command.TradingName,
		Website:        command.Website,
	}

	err = handler.messagePublisher.Publish(&event, context)

	if err != nil {
		slog.Error("Failed to publish message", "error", err)
		return nil, err
	}

	return organisationId, nil
}
