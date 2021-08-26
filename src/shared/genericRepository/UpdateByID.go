package genericRepository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func UpdateByID(collection string, id primitive.ObjectID, data interface{}, config *constants.TenantConfig) (primitive.ObjectID, error) {
	ctx := database.GetContext(30)

	var updateFields bson.D
	conv, _ := bson.Marshal(data)
	bson.Unmarshal(conv, &updateFields)
	update := bson.M{
		"$set": updateFields,
	}

	result, err := database.Collection(config.CollectionPrefix+collection).UpdateByID(ctx, id, update)

	if err != nil {
		return primitive.NilObjectID, err
	}

	upsertedId, _ := result.UpsertedID.(primitive.ObjectID)

	return upsertedId, err
}
