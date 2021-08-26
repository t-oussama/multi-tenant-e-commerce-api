package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/models"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/genericRepository"
)

func UpdateList(id primitive.ObjectID, data *models.UpdateProductsListInput, config *constants.TenantConfig) error {

	_, err := genericRepository.UpdateByID(constants.CollectionNames.PRODUCT_LISTS, id, data, config)

	return err
}
