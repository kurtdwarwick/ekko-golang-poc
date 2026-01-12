package repositories

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/core/data/access"
	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"

	"github.com/ekko-earth/shared/adapters"
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

func (repository *OrganisationRepository) GetOrganisationById(
	id string,
	transaction adapters.Transaction,
	ctx context.Context,
) (*entities.Organisation, error) {
	return repository.organisationDao.GetById(id, transaction, ctx)
}

func (repository *OrganisationRepository) GetAll(
	page *int32,
	size *int32,
	transaction adapters.Transaction,
	ctx context.Context,
) ([]entities.Organisation, error) {
	return repository.organisationDao.GetAll(page, size, transaction, ctx)
}
