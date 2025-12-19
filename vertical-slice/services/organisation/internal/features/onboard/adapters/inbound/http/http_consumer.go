package http

import (
	organisationCommands "organisation/internal/features/onboard/core/commands"
	organisationCommandHandlers "organisation/internal/features/onboard/core/commands/handlers"

	"encoding/json"

	"commands"
	"organisation/internal/adapters"

	"net/http"

	"github.com/google/uuid"
)

type OnboardOrganisationHttpConsumer struct {
}

func NewOnboardOrganisationHttpConsumer(
	server adapters.HttpServer,
	onboardOrganisationCommandHandler *organisationCommandHandlers.OnboardOrganisationCommandHandler,
) *OnboardOrganisationHttpConsumer {
	organisationsRoute := server.Router.PathPrefix("/organisations/onboard").Subrouter()

	organisationsRoute.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		onboardOrganisationHttpHandler(writer, request, onboardOrganisationCommandHandler)
	}).Methods("POST")

	return &OnboardOrganisationHttpConsumer{}
}

type OnboardOrganisationHttpDto struct {
	LegalName   string  `json:"legalName"`
	TradingName string  `json:"tradingName"`
	Website     *string `json:"website"`
}

func onboardOrganisationHttpHandler(
	writer http.ResponseWriter,
	request *http.Request,
	onboardOrganisationCommandHandler *organisationCommandHandlers.OnboardOrganisationCommandHandler,
) {
	var onboardOrganisationHttpDto OnboardOrganisationHttpDto

	err := json.NewDecoder(request.Body).Decode(&onboardOrganisationHttpDto)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	organisationId, err := onboardOrganisationCommandHandler.Handle(organisationCommands.OnboardOrganisationCommand{
		Command: commands.Command{
			ConversationId: uuid.New().String(),
		},
		LegalName:   onboardOrganisationHttpDto.LegalName,
		TradingName: onboardOrganisationHttpDto.TradingName,
		Website:     onboardOrganisationHttpDto.Website,
	})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(writer).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(map[string]string{
		"id": *organisationId,
	})
}
