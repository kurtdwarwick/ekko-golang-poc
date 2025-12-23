package repositories

import (
	"context"
	"errors"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/access"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/policies"

	"github.com/google/uuid"
)

type OrganisationRepository struct {
	organisationDao access.OrganisationDAO
	policyHandler   policies.PolicyHandler
}

func NewOrganisationRepository(
	organisationDao access.OrganisationDAO,
	policyHandler policies.PolicyHandler,
) *OrganisationRepository {
	return &OrganisationRepository{
		organisationDao: organisationDao,
		policyHandler:   policyHandler,
	}
}

func ValidateUniqueness(
	organisation entities.Organisation,
	organisationDao access.OrganisationDAO,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	count, err := organisationDao.Count(&entities.Organisation{LegalName: organisation.LegalName}, transaction, ctx)

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("organisation already exists")
	}

	return nil
}

func (repository *OrganisationRepository) OnboardOrganisation(
	organisation entities.Organisation,
	transaction adapters.Transaction,
	ctx context.Context,
) (*uuid.UUID, error) {
	err := ValidateUniqueness(organisation, repository.organisationDao, transaction, ctx)

	if err != nil {
		return nil, err
	}

	organisationId := uuid.New()

	err = repository.policyHandler.Apply(organisation)

	if err != nil {
		return nil, err
	}

	organisation.Id = organisationId

	err = repository.organisationDao.Create(&organisation, transaction, ctx)

	return &organisationId, err
}
