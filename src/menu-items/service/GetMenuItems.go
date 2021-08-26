package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func GetMenuItems(config *constants.TenantConfig) ([]entities.Category, error) {
	result := make([]entities.Category, 0)
	err := genericRepository.Find(constants.CollectionNames.CATEGORIES, bson.D{}, &result, config)
	return result, err
}
