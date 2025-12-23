package access

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
	"github.com/ekko-earth/shared/adapters"
)

type OrganisationDAO interface {
	Create(
		organisation *entities.Organisation,
		transaction adapters.Transaction,
		ctx context.Context,
	) error
	Count(
		organisation *entities.Organisation,
		transaction adapters.Transaction,
		ctx context.Context,
	) (int32, error)
}
