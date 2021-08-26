package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/repository"
	"store.tn/api/src/shared/constants"
)

func AddProductToSubCategory(productId primitive.ObjectID, subCategoryId primitive.ObjectID, config *constants.TenantConfig) error {
	err := repository.AddProductToSubCategory(productId, subCategoryId, config)
	return err
}
