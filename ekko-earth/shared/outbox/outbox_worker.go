package outbox

import (
	"context"
	"log/slog"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/google/uuid"

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/observability"

	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"
)

type OutboxWorker struct {
	repository       *OutboxRepository
	unitOfWork       adapters.UnitOfWork
	messagePublisher messagingAdapters.MessagePublisher
	configuration    OutboxWorkerConfiguration
	lockReference    string
}

type OutboxWorkerConfiguration struct {
	MaxWorkers   int
	PollInterval time.Duration
	BatchSize    int
}

type OutboxUnpublishedMessage struct {
	Message OutboxMessage
	Context context.Context
}

func NewOutboxWorker(
	repository *OutboxRepository,
	unitOfWork adapters.UnitOfWork,
	messagePublisher messagingAdapters.MessagePublisher,
	configuration OutboxWorkerConfiguration,
) *OutboxWorker {
	lockReference, err := machineid.ProtectedID("outbox-lock")

	if err != nil {
		panic(err)
	}

	return &OutboxWorker{
		repository:       repository,
		unitOfWork:       unitOfWork,
		messagePublisher: messagePublisher,
		configuration:    configuration,
		lockReference:    lockReference,
	}
}

func (worker *OutboxWorker) Start(ctx context.Context) error {
	slog.Info("Starting OutboxWorker")

	processChannel := make(chan OutboxUnpublishedMessage, worker.configuration.MaxWorkers)

	for range worker.configuration.MaxWorkers {
		go worker.processMessage(processChannel, context.WithValue(ctx, "traceId", uuid.New().String()))
	}

	ticker := time.NewTicker(worker.configuration.PollInterval)

	go func() error {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()

				return ctx.Err()
			case <-ticker.C:
				err := worker.execute(processChannel, ctx)

				if err != nil {
					slog.Error("Failed to execute outbox worker", "error", err)
				}
			default:
				time.Sleep(worker.configuration.PollInterval)
			}
		}
	}()

	return nil
}

func (worker *OutboxWorker) Stop(ctx context.Context) error {
	slog.Info("Stopping OutboxWorker")
	return nil
}

func (worker *OutboxWorker) execute(channel chan OutboxUnpublishedMessage, ctx context.Context) error {
	unsentMessages, err := worker.unitOfWork.Execute(
		func(transaction adapters.Transaction, ctx context.Context) (any, error) {
			err := worker.repository.LockUnsentMessages(
				worker.lockReference,
				worker.configuration.BatchSize,
				transaction,
				ctx,
			)

			if err != nil {
				slog.Error("Failed to lock unsent messages", "error", err)
				return nil, err
			}

			messages, err := worker.repository.GetUnsentMessages(
				worker.lockReference,
				worker.configuration.BatchSize,
				transaction,
				ctx,
			)

			if err != nil {
				slog.Error("Failed to get unsent messages", "error", err)
				return nil, err
			}

			unsentMessages := make([]OutboxUnpublishedMessage, len(messages))

			for i, message := range messages {
				if observability.Instrument {
					ctx = observability.PropagateContext(ctx, message.Headers)
				}

				unsentMessages[i] = OutboxUnpublishedMessage{Message: *message, Context: ctx}
			}

			return unsentMessages, nil
		},
		ctx,
	)

	if err != nil {
		slog.Error("Failed to get unsent messages", "error", err)
		return err
	}

	worker.processUnsentMessages(channel, unsentMessages.([]OutboxUnpublishedMessage))

	return nil
}

func (worker *OutboxWorker) processUnsentMessages(
	channel chan OutboxUnpublishedMessage,
	unsentMessages []OutboxUnpublishedMessage,
) {
	for _, unsentMessage := range unsentMessages {
		channel <- unsentMessage
	}
}

func (worker *OutboxWorker) processMessage(channel chan OutboxUnpublishedMessage, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case unsentMessage := <-channel:
			err := worker.messagePublisher.Publish(
				unsentMessage.Message.Message,
				unsentMessage.Message.MessageType,
				unsentMessage.Message.Headers,
				unsentMessage.Context,
			)

			if err != nil {
				slog.Error("Failed to publish message", "error", err)
			}

			worker.repository.RemoveMessage(unsentMessage.Message.Id, nil, unsentMessage.Context)
		}
	}
}
