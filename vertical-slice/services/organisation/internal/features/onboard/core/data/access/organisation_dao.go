package access

import "organisation/internal/data/entities"

type OrganisationDAO interface {
	Create(organisation *entities.Organisation) error
	Count(organisation *entities.Organisation) (int32, error)
}
