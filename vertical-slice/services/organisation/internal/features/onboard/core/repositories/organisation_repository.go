package repositories

import (
	"errors"
	"organisation/internal/data/entities"
	"organisation/internal/features/onboard/core/data/access"

	"github.com/google/uuid"

	"policies"
)

type OrganisationRepository struct {
	organisationDao access.OrganisationDAO
	organisations   map[string]entities.Organisation

	policyHandler policies.PolicyHandler
}

func NewOrganisationRepository(
	organisationDao access.OrganisationDAO,
	policyHandler policies.PolicyHandler,
) *OrganisationRepository {
	return &OrganisationRepository{
		organisationDao: organisationDao,
		organisations:   make(map[string]entities.Organisation),
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

func (repository *OrganisationRepository) OnboardOrganisation(organisation entities.Organisation) (*string, error) {
	err := ValidateUniqueness(organisation, repository.organisationDao)

	if err != nil {
		return nil, err
	}

	organisationId := uuid.New().String()

	err = repository.policyHandler.Apply(organisation)

	if err != nil {
		return nil, err
	}

	organisation.Id = organisationId

	err = repository.organisationDao.Create(&organisation)

	return &organisationId, err
}
