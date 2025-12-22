package gorm

import "github.com/ekko-earth/shared/gorm/adapters"

type OrganisationModel struct {
	adapters.GormModel
}

func (OrganisationModel) TableName() string {
	return "organisations"
}
