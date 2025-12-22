package access

import "github.com/ekko-earth/impact/internal/organisation/core/data/entities"

type OrganisationDAO interface {
	Save(organisation *entities.Organisation) error
}
