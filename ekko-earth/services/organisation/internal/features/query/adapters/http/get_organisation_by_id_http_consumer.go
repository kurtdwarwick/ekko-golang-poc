package http

import (
	"context"
	"errors"

	"github.com/ekko-earth/organisation/internal/features/query/core/queries"
	"github.com/ekko-earth/organisation/internal/features/query/core/queries/handlers"
	"github.com/ekko-earth/shared/http/adapters"
)

type GetOrganisationByIdHttpConsumer struct {
	queryHandler handlers.GetOrganisationByIdQueryHandler
}

func NewGetOrganisationByIdHttpConsumer(
	server *adapters.HttpServer,
	queryHandler handlers.GetOrganisationByIdQueryHandler,
) *GetOrganisationByIdHttpConsumer {
	consumer := &GetOrganisationByIdHttpConsumer{
		queryHandler: queryHandler,
	}

	adapters.NewHttpConsumer(
		server,
		adapters.HttpConsumerConfiguration{Route: "/organisations/{id}", Methods: []string{"GET"}},
		consumer,
	)

	return consumer
}

func (consumer *GetOrganisationByIdHttpConsumer) Consume(
	vars map[string]string,
	body GetOrganisationByIdHttpIncomingDto,
	ctx context.Context,
) (*GetOrganisationByIdHttpOutgoingDto, error) {
	id, ok := vars["id"]

	if !ok {
		return nil, errors.New("id is required")
	}

	organisation, err := consumer.queryHandler.Handle(queries.GetOrganisationByIdQuery{
		Id: id,
	}, ctx)

	if err != nil {
		return nil, err
	}

	return &GetOrganisationByIdHttpOutgoingDto{
		Id:          organisation.Id.String(),
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	}, nil
}
