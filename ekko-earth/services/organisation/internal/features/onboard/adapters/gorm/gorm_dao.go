package gorm

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/adapters/gorm/models"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"

	"gorm.io/gorm"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
)

type GormOrganisationDAO struct {
	db *gorm.DB
}

func NewGormOrganizationDAO(database gormAdapters.GormDatabase) *GormOrganisationDAO {
	database.Database.AutoMigrate(&models.OrganisationModel{})

	return &GormOrganisationDAO{db: database.Database}
}

func (dao *GormOrganisationDAO) Create(organisation *entities.Organisation) error {
	context := context.TODO()

	err := dao.db.Transaction(func(tx *gorm.DB) error {
		err := gorm.G[models.OrganisationModel](dao.db).Create(context, &models.OrganisationModel{
			GormModel: gormAdapters.GormModel{
				Id: organisation.Id,
			},
			LegalName:   organisation.LegalName,
			TradingName: organisation.TradingName,
			Website:     organisation.Website,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (dao *GormOrganisationDAO) Count(organisation *entities.Organisation) (int32, error) {
	var count int64

	err := dao.db.Model(&models.OrganisationModel{}).Where(&models.OrganisationModel{
		GormModel: gormAdapters.GormModel{
			Id: organisation.Id,
		},
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	}).Count(&count).Error

	return int32(count), err
}
