package mongodb

import (
	"context"
	"log/slog"

	"github.com/ekko-earth/impact/internal/organisation/core/data/entities"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	mongoAdapters "github.com/ekko-earth/shared/mongodb/adapters"
)

type MongoDBOrganisationDAO struct {
	database   *mongoAdapters.MongoDatabase
	collection *mongo.Collection
}

func NewMongoDBOrganisationDAO(database mongoAdapters.MongoDatabase) *MongoDBOrganisationDAO {
	err := database.Client.Database(database.Configuration.Database).CreateCollection(context.TODO(), "organisations")

	if err != nil {
		panic(err)
	}

	collection := database.Client.Database(database.Configuration.Database).Collection("organisations")

	return &MongoDBOrganisationDAO{
		database:   &database,
		collection: collection,
	}
}

func (dao *MongoDBOrganisationDAO) Save(organisation *entities.Organisation, context context.Context) error {
	options := options.UpdateOne().SetUpsert(true)

	model := OrganisationModel{
		MongoModel: mongoAdapters.MongoModel{
			Id: organisation.Id,
		},
	}

	_, err := dao.collection.UpdateOne(
		context,
		bson.M{"_id": organisation.Id},
		bson.M{"$set": model},
		options,
	)

	if err != nil {
		slog.Error("Failed to save organisation", "error", err)
		return err
	}

	return nil
}
