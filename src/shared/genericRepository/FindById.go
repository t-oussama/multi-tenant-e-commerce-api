package genericRepository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func FindById(collection string, id primitive.ObjectID, result interface{}, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)
	err := database.Collection(config.CollectionPrefix+collection).FindOne(ctx, bson.D{{"_id", id}}).Decode(result)

	return err
}
