package access

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/core/data/entities"
	"github.com/ekko-earth/shared/adapters"
)

type OrganisationDAO interface {
	GetById(
		id string,
		transaction adapters.Transaction,
		ctx context.Context,
	) (*entities.Organisation, error)

	GetAll(
		page *int32,
		size *int32,
		transaction adapters.Transaction,
		ctx context.Context,
	) ([]entities.Organisation, error)
}
