package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func GetAll(config *constants.TenantConfig) ([]entities.Product, error) {
	result := make([]entities.Product, 0)
	err := genericRepository.Find(constants.CollectionNames.PRODUCTS, bson.D{}, &result, config)
	return result, err
}
