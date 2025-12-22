package adapters

import (
	"github.com/google/uuid"
)

type MongoModel struct {
	Id uuid.UUID `bson:"_id,omitempty"`
}
