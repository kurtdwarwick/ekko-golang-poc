package mongodb

import (
	"github.com/ekko-earth/shared/mongodb/adapters"
)

type OrganisationModel struct {
	adapters.MongoModel `bson:",inline"`
}
