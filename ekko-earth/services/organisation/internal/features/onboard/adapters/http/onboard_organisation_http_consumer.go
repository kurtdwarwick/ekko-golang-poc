package http

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/commands"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/commands/handlers"
	"github.com/ekko-earth/shared/http/adapters"
)

type OnboardOrganisationHttpConsumer struct {
	onboardOrganisationCommandHandler *handlers.OnboardOrganisationCommandHandler
}

func NewOnboardOrganisationHttpConsumer(
	server *adapters.HttpServer,
	onboardOrganisationCommandHandler *handlers.OnboardOrganisationCommandHandler,
) *OnboardOrganisationHttpConsumer {
	consumer := &OnboardOrganisationHttpConsumer{
		onboardOrganisationCommandHandler: onboardOrganisationCommandHandler,
	}

	adapters.NewHttpConsumer(
		server,
		adapters.HttpConsumerConfiguration{Route: "/organisations/onboard", Methods: []string{"POST"}},
		consumer,
	)

	return consumer
}

func (consumer *OnboardOrganisationHttpConsumer) Consume(
	vars map[string]string,
	body OnboardOrganisationHttpIncomingDto,
	ctx context.Context,
) (*OnboardOrganisationHttpOutgoingDto, error) {
	command := commands.OnboardOrganisationCommand{
		LegalName:   body.LegalName,
		TradingName: body.TradingName,
		Website:     &body.Website,
	}

	organisationId, err := consumer.onboardOrganisationCommandHandler.Handle(command, ctx)

	if err != nil {
		return nil, err
	}

	return &OnboardOrganisationHttpOutgoingDto{
		Id: organisationId.String(),
	}, nil
}
