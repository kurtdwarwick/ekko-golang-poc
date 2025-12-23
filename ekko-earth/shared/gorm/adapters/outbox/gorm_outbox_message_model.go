package outbox

import (
	"time"

	"github.com/ekko-earth/shared/gorm/adapters"
	"gorm.io/datatypes"
)

type GormOutboxMessageModel struct {
	adapters.GormModel

	MessageType string         `gorm:"index"`
	Message     datatypes.JSON `gorm:"type:jsonb"`

	LockedAt      *time.Time `gorm:"null;"`
	LockReference *string    `gorm:"null; index"`
}

func (GormOutboxMessageModel) TableName() string {
	return "outbox"
}
