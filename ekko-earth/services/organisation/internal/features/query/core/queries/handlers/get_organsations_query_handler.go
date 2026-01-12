package handlers

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"
	"github.com/ekko-earth/organisation/internal/features/query/core/queries"
	"github.com/ekko-earth/organisation/internal/features/query/core/repositories"
)

type GetOrganisationsQueryHandler struct {
	repository *repositories.OrganisationRepository
}

func NewGetOrganisationsQueryHandler(
	repository *repositories.OrganisationRepository,
) *GetOrganisationsQueryHandler {
	return &GetOrganisationsQueryHandler{
		repository: repository,
	}
}

func (handler *GetOrganisationsQueryHandler) Handle(
	query queries.GetOrganisationsQuery,
	ctx context.Context,
) ([]entities.Organisation, error) {
	return handler.repository.GetAll(query.Page, query.Size, nil, ctx)
}
