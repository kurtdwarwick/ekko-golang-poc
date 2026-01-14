package gorm

import (
	"time"

	"gorm.io/datatypes"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
)

type GormOutboxMessageModel struct {
	gormAdapters.GormModel

	MessageType   string         `gorm:"index"`
	Message       datatypes.JSON `gorm:"type:jsonb"`
	Headers       datatypes.JSON `gorm:"type:jsonb"`
	LockedAt      *time.Time     `gorm:"null;"`
	LockReference *string        `gorm:"null; index"`
}

func (GormOutboxMessageModel) TableName() string {
	return "outbox"
}
