package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
	"store.tn/api/src/shared/entities"
)

func CreateCategory(data entities.Category, config *constants.TenantConfig) (primitive.ObjectID, error) {
	ctx := database.GetContext(30)

	result, err := database.Collection(config.CollectionPrefix+constants.CollectionNames.CATEGORIES).InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	return id, err
}
