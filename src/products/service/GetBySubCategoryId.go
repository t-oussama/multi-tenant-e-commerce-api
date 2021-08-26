package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func GetBySubCategoryId(subCategoryId primitive.ObjectID, config *constants.TenantConfig) (*[]entities.Product, error) {
	results := new([]entities.Product)
	err := repository.GetBySubCategoryId(subCategoryId, &results, config)
	return results, err
}
