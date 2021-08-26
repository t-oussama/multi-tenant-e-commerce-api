package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func GetLists(config *constants.TenantConfig) ([]entities.ProductList, error) {
	result := make([]entities.ProductList, 0)
	err := genericRepository.Find(constants.CollectionNames.PRODUCT_LISTS, bson.D{}, &result, config)
	return result, err
}
