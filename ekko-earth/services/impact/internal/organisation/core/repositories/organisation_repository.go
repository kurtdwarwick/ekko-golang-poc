package repositories

import (
	"github.com/ekko-earth/impact/internal/organisation/core/data/access"
	"github.com/ekko-earth/impact/internal/organisation/core/data/entities"
)

type OrganisationRepository struct {
	organisationDao access.OrganisationDAO
}

func NewOrganisationRepository(
	organisationDao access.OrganisationDAO,
) *OrganisationRepository {
	return &OrganisationRepository{
		organisationDao: organisationDao,
	}
}

func (repository *OrganisationRepository) OnboardOrganisation(organisation entities.Organisation) error {
	err := repository.organisationDao.Save(&organisation)

	return err
}
