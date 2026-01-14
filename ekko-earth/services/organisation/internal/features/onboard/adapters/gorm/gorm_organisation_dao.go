package gorm

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/adapters/gorm/models"
	"github.com/ekko-earth/organisation/internal/features/onboard/core"
	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

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
	organisation *core.Organisation,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	spanContext, span := observability.Tracer.Start(ctx, "GormOrganisationDAO.Create")

	defer span.End()

	database := dao.database.Database

	if transaction != nil {
		database = transaction.(*gormAdapters.GormTransaction).Transaction
	}

	err := gorm.G[models.GormOrganisationModel](database).Create(spanContext, &models.GormOrganisationModel{
		GormModel: gormAdapters.GormModel{
			Id: organisation.Id,
		},
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	})

	span.AddEvent(
		"Organisation created",
		trace.WithAttributes(attribute.String("organisation.id", organisation.Id.String())),
	)

	return err
}

func (dao *GormOrganisationDAO) Count(
	organisation *core.Organisation,
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
