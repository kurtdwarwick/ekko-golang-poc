package outbox

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/outbox"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
)

type GormOutboxDAO struct {
	database gormAdapters.GormDatabase
}

func NewGormOutboxDAO(database gormAdapters.GormDatabase) *GormOutboxDAO {
	database.Database.AutoMigrate(&GormOutboxMessageModel{})

	return &GormOutboxDAO{database: database}
}

func (dao *GormOutboxDAO) Create(
	outboxMessage *outbox.OutboxMessage,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	message, err := json.Marshal(outboxMessage.Message)

	if err != nil {
		return err
	}

	return gorm.G[GormOutboxMessageModel](database).Create(ctx, &GormOutboxMessageModel{
		MessageType: outboxMessage.GetMessageType(),
		Message:     datatypes.JSON(message),
	})
}

func (dao *GormOutboxDAO) Update(
	where map[string]any,
	update *outbox.OutboxMessage,
	limit int,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	_, err := gorm.G[GormOutboxMessageModel](database).Where(where).Limit(limit).Updates(ctx, GormOutboxMessageModel{
		LockedAt:      update.LockedAt,
		LockReference: update.LockReference,
	})

	return err
}

func (dao *GormOutboxDAO) Find(
	where map[string]any,
	limit int,
	transaction adapters.Transaction,
	ctx context.Context,
) ([]outbox.OutboxMessage, error) {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	outboxMessagesModels, err := gorm.G[GormOutboxMessageModel](database).Where(where).Limit(limit).Find(ctx)

	if err != nil {
		return nil, err
	}

	outboxMessages := make([]outbox.OutboxMessage, len(outboxMessagesModels))

	for i, outboxMessageModel := range outboxMessagesModels {
		outboxMessages[i] = outbox.OutboxMessage{
			Id:            outboxMessageModel.Id,
			MessageType:   outboxMessageModel.MessageType,
			Message:       outboxMessageModel.Message,
			CreatedAt:     outboxMessageModel.CreatedAt,
			LockedAt:      outboxMessageModel.LockedAt,
			LockReference: outboxMessageModel.LockReference,
		}
	}

	return outboxMessages, nil
}

func (dao *GormOutboxDAO) Delete(
	id uuid.UUID,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	_, err := gorm.G[GormOutboxMessageModel](database).Where("id = ?", id).Delete(ctx)

	return err
}
