package gorm

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/adapters/gorm/models"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
	"github.com/ekko-earth/shared/adapters"

	"gorm.io/gorm"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
)

type GormOrganisationDAO struct {
	database gormAdapters.GormDatabase
}

func NewGormOrganizationDAO(database gormAdapters.GormDatabase) *GormOrganisationDAO {
	database.Database.AutoMigrate(&models.GormOrganisationModel{})

	return &GormOrganisationDAO{database: database}
}

func (dao *GormOrganisationDAO) Create(
	organisation *entities.Organisation,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	err := gorm.G[models.GormOrganisationModel](database).Create(ctx, &models.GormOrganisationModel{
		GormModel: gormAdapters.GormModel{
			Id: organisation.Id,
		},
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	})

	return err
}

func (dao *GormOrganisationDAO) Count(
	organisation *entities.Organisation,
	transaction adapters.Transaction,
	ctx context.Context,
) (int32, error) {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	var count int64

	err := database.Model(&models.GormOrganisationModel{}).Where(&models.GormOrganisationModel{
		GormModel: gormAdapters.GormModel{
			Id: organisation.Id,
		},
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	}).Count(&count).Error

	return int32(count), err
}
