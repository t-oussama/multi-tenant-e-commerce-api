package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
	"store.tn/api/src/shared/models"
)

func GetProductsByIdsList(ids []primitive.ObjectID, results *[]models.ProductBaseInfo, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)

	query := bson.M{"_id": bson.M{"$in": ids}}

	cur, err := database.Collection(config.CollectionPrefix+constants.CollectionNames.PRODUCTS).Find(ctx, query)

	if err != nil {
		return err
	}

	if err = cur.All(ctx, results); err != nil {
		return err
	}

	return err
}
