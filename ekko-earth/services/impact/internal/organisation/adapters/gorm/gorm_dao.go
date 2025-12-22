package gorm

import (
	"github.com/ekko-earth/impact/internal/organisation/core/data/entities"

	"gorm.io/gorm"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
)

type GormOrganisationDAO struct {
	db *gorm.DB
}

func NewGormOrganizationDAO(database gormAdapters.GormDatabase) *GormOrganisationDAO {
	database.Database.AutoMigrate(&OrganisationModel{})

	return &GormOrganisationDAO{db: database.Database}
}

func (dao *GormOrganisationDAO) Save(organisation *entities.Organisation) error {
	err := dao.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(&OrganisationModel{
			GormModel: gormAdapters.GormModel{
				Id: organisation.Id,
			},
		})

		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	return err
}
