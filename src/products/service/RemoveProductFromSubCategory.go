package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/repository"
	"store.tn/api/src/shared/constants"
)

func RemoveProductFromSubCategory(productId primitive.ObjectID, subCategoryId primitive.ObjectID, config *constants.TenantConfig) error {
	err := repository.RemoveProductFromSubCategory(productId, subCategoryId, config)
	return err
}
