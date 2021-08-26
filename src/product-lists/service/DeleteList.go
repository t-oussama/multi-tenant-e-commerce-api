package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/genericRepository"
)

func DeleteList(id primitive.ObjectID, config *constants.TenantConfig) error {

	err := genericRepository.DeleteById(constants.CollectionNames.PRODUCT_LISTS, id, config)

	return err
}
