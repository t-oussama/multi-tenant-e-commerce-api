package genericRepository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func InsertOne(collection string, data interface{}, config *constants.TenantConfig) (primitive.ObjectID, error) {
	ctx := database.GetContext(30)
	result, insertErr := database.Collection(config.CollectionPrefix+collection).InsertOne(ctx, data)

	id, _ := result.InsertedID.(primitive.ObjectID)

	return id, insertErr
}
