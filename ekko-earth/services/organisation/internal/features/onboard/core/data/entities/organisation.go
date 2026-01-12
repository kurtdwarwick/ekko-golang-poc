package entities

import "github.com/google/uuid"

type Organisation struct {
	Id uuid.UUID

	LegalName   string
	TradingName string

	Website *string
}
