package handlers

import (
	"log/slog"
	"organisation/internal/data/entities"
	"organisation/internal/features/onboard/core/commands"
	"organisation/internal/features/onboard/core/repositories"
)

type OnboardOrganisationCommandHandler struct {
	repository *repositories.OrganisationRepository
}

func NewOnboardOrganisationCommandHandler(
	repository *repositories.OrganisationRepository,
) *OnboardOrganisationCommandHandler {
	return &OnboardOrganisationCommandHandler{
		repository: repository,
	}
}

func (handler *OnboardOrganisationCommandHandler) Handle(command commands.OnboardOrganisationCommand) (*string, error) {
	organisationId, err := handler.repository.OnboardOrganisation(
		entities.Organisation{
			LegalName:   command.LegalName,
			TradingName: command.TradingName,
			Website:     command.Website,
		})

	if err != nil {
		slog.Error("Failed to onboard organisation", "error", err)
		return nil, err
	}

	return organisationId, err
}
