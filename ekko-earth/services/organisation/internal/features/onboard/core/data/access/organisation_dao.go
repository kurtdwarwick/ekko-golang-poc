package access

import "github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"

type OrganisationDAO interface {
	Create(organisation *entities.Organisation) error
	Count(organisation *entities.Organisation) (int32, error)
}
