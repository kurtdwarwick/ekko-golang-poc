package messaging

import (
	"reflect"

	"github.com/google/uuid"
)

type Message struct {
	ConversationId uuid.UUID
}

type WithMessageType interface {
	GetMessageType() string
}

func (message *Message) GetMessageType() string {
	return reflect.TypeOf(message).Name()
}
