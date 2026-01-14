package core

import (
	"context"

	"github.com/ekko-earth/shared/adapters"
)

type OrganisationDAO interface {
	Create(
		organisation *Organisation,
		transaction adapters.Transaction,
		ctx context.Context,
	) error
	Count(
		organisation *Organisation,
		transaction adapters.Transaction,
		ctx context.Context,
	) (int32, error)
}
