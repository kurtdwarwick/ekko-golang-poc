package gorm

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/adapters/gorm/models"
	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"
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

func (dao *GormOrganisationDAO) GetById(
	id string,
	transaction adapters.Transaction,
	ctx context.Context,
) (*entities.Organisation, error) {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	organisation, err := gorm.G[models.GormOrganisationModel](database).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	return &entities.Organisation{
		Id:          organisation.Id,
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	}, nil
}

func (dao *GormOrganisationDAO) GetAll(
	page *int32,
	size *int32,
	transaction adapters.Transaction,
	ctx context.Context,
) ([]entities.Organisation, error) {
	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	organisationModels, err := gorm.G[models.GormOrganisationModel](
		database,
	).Offset(int(*page)).
		Limit(int(*size)).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	organisations := make([]entities.Organisation, 0, len(organisationModels))

	for _, organisationModel := range organisationModels {
		organisations = append(organisations, entities.Organisation{
			Id:          organisationModel.Id,
			LegalName:   organisationModel.LegalName,
			TradingName: organisationModel.TradingName,
			Website:     organisationModel.Website,
		})
	}

	return organisations, nil
}
