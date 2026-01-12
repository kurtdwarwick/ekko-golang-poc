package handlers

import (
	"context"
	"log/slog"

	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"
	"github.com/ekko-earth/organisation/internal/features/query/core/queries"
	"github.com/ekko-earth/organisation/internal/features/query/core/repositories"
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
	slog.Info("Getting organisation by ID", "ID", query.Id)
	return handler.repository.GetOrganisationById(query.Id, nil, ctx)
}
