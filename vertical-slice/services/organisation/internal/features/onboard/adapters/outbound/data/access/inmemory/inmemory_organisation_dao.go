package inmemory

import (
	"organisation/internal/data/entities"
)

type InMemoryOrganisationDAO struct {
	organisations map[string]entities.Organisation
}

func NewInMemoryOrganisationDAO() *InMemoryOrganisationDAO {
	return &InMemoryOrganisationDAO{
		organisations: make(map[string]entities.Organisation),
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
