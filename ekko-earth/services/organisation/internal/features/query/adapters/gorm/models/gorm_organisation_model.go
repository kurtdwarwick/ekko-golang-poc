package models

import "github.com/ekko-earth/shared/gorm/adapters"

type GormOrganisationModel struct {
	adapters.GormModel

	LegalName   string
	TradingName string
	Website     *string
}

func (GormOrganisationModel) TableName() string {
	return "organisations"
}
