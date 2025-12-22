package repositories

import (
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

func ValidateUniqueness(organisation entities.Organisation, organisationDao access.OrganisationDAO) error {
	count, err := organisationDao.Count(&entities.Organisation{LegalName: organisation.LegalName})

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("organisation already exists")
	}

	return nil
}

func (repository *OrganisationRepository) OnboardOrganisation(organisation entities.Organisation) (*uuid.UUID, error) {
	err := ValidateUniqueness(organisation, repository.organisationDao)

	if err != nil {
		return nil, err
	}

	organisationId := uuid.New()

	err = repository.policyHandler.Apply(organisation)

	if err != nil {
		return nil, err
	}

	organisation.Id = organisationId

	err = repository.organisationDao.Create(&organisation)

	return &organisationId, err
}
