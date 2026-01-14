package handlers

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"
	"github.com/ekko-earth/organisation/internal/features/query/core/queries"
	"github.com/ekko-earth/organisation/internal/features/query/core/repositories"
	"github.com/ekko-earth/shared/observability"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type GetOrganisationByIdQueryHandler struct {
	repository *repositories.OrganisationRepository
}

func NewGetOrganisationByIdQueryHandler(
	repository *repositories.OrganisationRepository,
) *GetOrganisationByIdQueryHandler {
	return &GetOrganisationByIdQueryHandler{
		repository: repository,
	}
}

func (handler *GetOrganisationByIdQueryHandler) Handle(
	query queries.GetOrganisationByIdQuery,
	ctx context.Context,
) (*entities.Organisation, error) {
	spanContext, span := observability.Tracer.Start(ctx, "GetOrganisationByIdQueryHandler.Handle")

	defer span.End()

	span.SetAttributes(attribute.String("query.id", query.Id))

	organisation, err := handler.repository.GetOrganisationById(query.Id, nil, spanContext)

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent(
		"Organisation retrieved",
		trace.WithAttributes(attribute.String("organisation.id", organisation.Id.String())),
	)

	return organisation, nil
}
