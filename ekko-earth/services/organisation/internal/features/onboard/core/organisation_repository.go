package core

import (
	"context"
	"errors"

	"github.com/ekko-earth/shared/adapters"
	"github.com/ekko-earth/shared/policies"

	"github.com/ekko-earth/shared/observability"
	"github.com/google/uuid"
)

type OrganisationRepository struct {
	organisationDao OrganisationDAO
	policyHandler   policies.PolicyHandler
}

func NewOrganisationRepository(
	organisationDao OrganisationDAO,
	policyHandler policies.PolicyHandler,
) *OrganisationRepository {
	return &OrganisationRepository{
		organisationDao: organisationDao,
		policyHandler:   policyHandler,
	}
}

func ValidateUniqueness(
	organisation Organisation,
	organisationDao OrganisationDAO,
	transaction adapters.Transaction,
	ctx context.Context,
) error {
	count, err := organisationDao.Count(&Organisation{LegalName: organisation.LegalName}, transaction, ctx)

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("organisation already exists")
	}

	return nil
}

func (repository *OrganisationRepository) OnboardOrganisation(
	organisation Organisation,
	transaction adapters.Transaction,
	ctx context.Context,
) (*uuid.UUID, error) {
	spanContext, span := observability.Tracer.Start(ctx, "OrganisationRepository.OnboardOrganisation")

	defer span.End()

	err := ValidateUniqueness(organisation, repository.organisationDao, transaction, spanContext)

	if err != nil {
		return nil, err
	}

	organisationId := uuid.New()

	err = repository.policyHandler.Apply(organisation)

	if err != nil {
		return nil, err
	}

	organisation.Id = organisationId

	err = repository.organisationDao.Create(&organisation, transaction, spanContext)

	return &organisationId, err
}
