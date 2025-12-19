package entities

import "data"

type Organisation struct {
	data.Entity

	LegalName   string
	TradingName string
	Website     *string
}
