package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
	"store.tn/api/src/shared/entities"
)

func InitializeSubCategoryProducts(id primitive.ObjectID, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)
	data := entities.SubCategoryProducts{
		Id:       id,
		Products: make([]primitive.ObjectID, 0),
	}

	_, err := database.Collection(config.CollectionPrefix+constants.CollectionNames.SUB_CATEGORY_PRODUCTS).InsertOne(ctx, data)

	return err
}
