package genericRepository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func DeleteById(collection string, id primitive.ObjectID, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)
	_, err := database.Collection(config.CollectionPrefix+collection).DeleteOne(ctx, bson.D{{"_id", id}})

	return err
}
