package inmemory

import (
	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
	"github.com/google/uuid"
)

type InMemoryOrganisationDAO struct {
	organisations map[uuid.UUID]entities.Organisation
}

func NewInMemoryOrganisationDAO() *InMemoryOrganisationDAO {
	return &InMemoryOrganisationDAO{
		organisations: make(map[uuid.UUID]entities.Organisation),
	}
}

func (dao *InMemoryOrganisationDAO) Create(
	organisation *entities.Organisation) error {
	dao.organisations[organisation.Id] = *organisation

	return nil
}

func (dao *InMemoryOrganisationDAO) Count(organisation *entities.Organisation) (int32, error) {
	return 0, nil
}
