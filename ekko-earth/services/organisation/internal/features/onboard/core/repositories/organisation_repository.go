package repositories

import (
	"context"
	"errors"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/access"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"

	"github.com/google/uuid"

	"github.com/ekko-earth/shared/policies"
)

type OrganisationRepository struct {
	organisationDao access.OrganisationDAO

	policyHandler policies.PolicyHandler
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
	context context.Context,
) error {
	count, err := organisationDao.Count(&entities.Organisation{LegalName: organisation.LegalName}, context)

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
	context context.Context,
) (*uuid.UUID, error) {
	err := ValidateUniqueness(organisation, repository.organisationDao, context)

	if err != nil {
		return nil, err
	}

	organisationId := uuid.New()

	err = repository.policyHandler.Apply(organisation)

	if err != nil {
		return nil, err
	}

	organisation.Id = organisationId

	err = repository.organisationDao.Create(&organisation, context)

	return &organisationId, err
}
