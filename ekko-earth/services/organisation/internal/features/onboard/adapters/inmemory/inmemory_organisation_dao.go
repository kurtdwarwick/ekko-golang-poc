package inmemory

import (
	"github.com/ekko-earth/organisation/internal/features/onboard/core"
	"github.com/google/uuid"
)

type InMemoryOrganisationDAO struct {
	organisations map[uuid.UUID]core.Organisation
}

func NewInMemoryOrganisationDAO() *InMemoryOrganisationDAO {
	return &InMemoryOrganisationDAO{
		organisations: make(map[uuid.UUID]core.Organisation),
	}
}

func (dao *InMemoryOrganisationDAO) Create(
	organisation *core.Organisation) error {
	dao.organisations[organisation.Id] = *organisation

	return nil
}

func (dao *InMemoryOrganisationDAO) Count(organisation *core.Organisation) (int32, error) {
	return 0, nil
}
