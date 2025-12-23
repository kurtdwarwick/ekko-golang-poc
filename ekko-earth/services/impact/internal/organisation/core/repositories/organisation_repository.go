package repositories

import (
	"context"

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

func (repository *OrganisationRepository) OnboardOrganisation(
	organisation entities.Organisation,
	ctx context.Context,
) error {
	err := repository.organisationDao.Save(&organisation, ctx)

	return err
}
