package handlers

import (
	"github.com/ekko-earth/impact/internal/organisation/core/data/entities"
	"github.com/ekko-earth/impact/internal/organisation/core/events"
	"github.com/ekko-earth/impact/internal/organisation/core/repositories"
)

type OrganisationOnboardedEventHandler struct {
	repository *repositories.OrganisationRepository
}

func NewOrganisationOnboardedEventHandler(
	repository *repositories.OrganisationRepository,
) *OrganisationOnboardedEventHandler {
	return &OrganisationOnboardedEventHandler{
		repository: repository,
	}
}

func (handler *OrganisationOnboardedEventHandler) Handle(message events.OrganisationOnboardedEvent) (any, error) {
	handler.repository.OnboardOrganisation(entities.Organisation{
		Id: message.OrganisationId,
	})

	return message.OrganisationId, nil
}
