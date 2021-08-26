package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"store.tn/api/src/products/models"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
	"store.tn/api/src/shared/entities"
)

func GetBySubCategoryId(subCategoryId primitive.ObjectID, results **[]entities.Product, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)

	var result []models.PopulatedSubCategoryProducts

	matchStage := bson.D{{"$match", bson.D{{"_id", subCategoryId}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "products"}, {"localField", "products"}, {"foreignField", "_id"}, {"as", "products"}}}}

	collection := database.Collection(config.CollectionPrefix + constants.CollectionNames.SUB_CATEGORY_PRODUCTS)
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return err
	}

	*results = &result[0].Products
	return err
}
