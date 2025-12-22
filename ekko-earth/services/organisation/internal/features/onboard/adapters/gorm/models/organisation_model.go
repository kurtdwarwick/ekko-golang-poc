package models

import "github.com/ekko-earth/shared/gorm/adapters"

type OrganisationModel struct {
	adapters.GormModel

	LegalName   string
	TradingName string
	Website     *string
}

func (OrganisationModel) TableName() string {
	return "organisations"
}
