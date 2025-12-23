package adapters

import (
	"time"

	"github.com/google/uuid"
)

type GormModel struct {
	Id uuid.UUID `gorm:"type:uuid; default:gen_random_uuid(); primary_key"`

	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"null;"`
}
