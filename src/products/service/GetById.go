package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func GetById(id primitive.ObjectID, config *constants.TenantConfig) (*entities.Product, error) {
	var result *entities.Product = new(entities.Product)
	err := genericRepository.FindById(constants.CollectionNames.PRODUCTS, id, result, config)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return result, err
}
