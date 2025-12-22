package adapters

import "github.com/google/uuid"

type GormModel struct {
	Id uuid.UUID `gorm:"type:uuid; primary_key"`
}
