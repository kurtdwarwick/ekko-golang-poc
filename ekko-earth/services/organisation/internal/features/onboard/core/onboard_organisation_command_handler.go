package core

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"
	"github.com/ekko-earth/shared/observability"

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/messaging"
	"github.com/ekko-earth/shared/outbox"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OnboardOrganisationCommandHandler struct {
	repository       *OrganisationRepository
	unitOfWork       adapters.UnitOfWork
	messagePublisher messagingAdapters.MessagePublisher
	outboxRepository *outbox.OutboxRepository
}

func NewOnboardOrganisationCommandHandler(
	repository *OrganisationRepository,
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
	command OnboardOrganisationCommand,
	ctx context.Context,
) (*uuid.UUID, error) {
	spanContext, span := observability.Tracer.Start(ctx, "OnboardOrganisationCommandHandler.Handle")

	defer span.End()

	organisationId, err := handler.unitOfWork.Execute(
		func(transaction adapters.Transaction, ctx context.Context) (any, error) {
			organisationId, err := handler.repository.OnboardOrganisation(
				Organisation{
					LegalName:   command.LegalName,
					TradingName: command.TradingName,
					Website:     command.Website,
				}, transaction, ctx)

			if err != nil {
				slog.Error("Failed to onboard organisation", "error", err)
				return nil, err
			}

			event := &OrganisationOnboardedEvent{
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

			span.AddEvent(
				"Organisation onboarded",
				trace.WithAttributes(attribute.String("organisation.id", organisationId.String())),
			)

			err = handler.outboxRepository.ScheduleMessage(&outbox.OutboxMessage{
				MessageType: event.GetMessageType(),
				Message:     *event,
				Headers:     make(map[string]any),
			}, transaction, ctx)

			if err != nil {
				slog.Error("Failed to publish message", "error", err)
				return nil, err
			}

			return organisationId, nil
		},
		spanContext,
	)

	return organisationId.(*uuid.UUID), err
}
