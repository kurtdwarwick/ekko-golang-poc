package access

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
)

type OrganisationDAO interface {
	Create(organisation *entities.Organisation, context context.Context) error
	Count(organisation *entities.Organisation, context context.Context) (int32, error)
}
