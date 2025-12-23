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

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/messaging"
	"github.com/ekko-earth/shared/outbox"
)

type OnboardOrganisationCommandHandler struct {
	repository       *repositories.OrganisationRepository
	unitOfWork       adapters.UnitOfWork
	messagePublisher messagingAdapters.MessagePublisher
	outboxRepository *outbox.OutboxRepository
}

func NewOnboardOrganisationCommandHandler(
	repository *repositories.OrganisationRepository,
	unitOfWork adapters.UnitOfWork,
	messagePublisher messagingAdapters.MessagePublisher,
	outboxRepository *outbox.OutboxRepository,
) *OnboardOrganisationCommandHandler {
	return &OnboardOrganisationCommandHandler{
		repository:       repository,
		unitOfWork:       unitOfWork,
		messagePublisher: messagePublisher,
		outboxRepository: outboxRepository,
	}
}

func (handler *OnboardOrganisationCommandHandler) Handle(
	command commands.OnboardOrganisationCommand,
	ctx context.Context,
) (any, error) {
	organisationId, err := handler.unitOfWork.Execute(
		func(transaction adapters.Transaction, ctx context.Context) (any, error) {
			organisationId, err := handler.repository.OnboardOrganisation(
				entities.Organisation{
					LegalName:   command.LegalName,
					TradingName: command.TradingName,
					Website:     command.Website,
				}, transaction, ctx)

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

			err = handler.outboxRepository.ScheduleMessage(&outbox.OutboxMessage{
				MessageType: event.GetMessageType(),
				Message:     event,
			}, transaction, ctx)

			if err != nil {
				slog.Error("Failed to publish message", "error", err)
				return nil, err
			}

			return organisationId, nil
		},
		ctx,
	)

	return organisationId, err
}
