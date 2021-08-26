package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func AddProductToSubCategory(productId primitive.ObjectID, subCategoryId primitive.ObjectID, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)

	update := bson.M{
		"$push": bson.M{
			"products": productId,
		},
	}

	_, err := database.Collection(config.CollectionPrefix+constants.CollectionNames.SUB_CATEGORY_PRODUCTS).UpdateByID(ctx, subCategoryId, update)

	return err
}
