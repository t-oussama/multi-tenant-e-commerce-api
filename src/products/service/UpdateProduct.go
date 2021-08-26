package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func UpdateProduct(id primitive.ObjectID, data *entities.Product, config *constants.TenantConfig) error {

	_, err := genericRepository.UpdateByID(config.CollectionPrefix+constants.CollectionNames.PRODUCTS, id, data, config)

	return err
}
