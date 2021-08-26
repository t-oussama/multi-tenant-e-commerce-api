package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func RemoveProductsFromList(id primitive.ObjectID, ids []primitive.ObjectID, config *constants.TenantConfig) (*entities.ProductList, error) {

	productsList := new(entities.ProductList)

	findError := genericRepository.FindById(constants.CollectionNames.PRODUCT_LISTS, id, productsList, config)

	if findError != nil {
		return nil, findError
	}

	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(productsList.Products); j++ {
			if productsList.Products[j].Id == ids[i] {
				productsList.Products = append(productsList.Products[:j], productsList.Products[j+1:]...)
			}
		}
	}

	id, err := genericRepository.UpdateByID(constants.CollectionNames.PRODUCT_LISTS, id, *productsList, config)

	if err != nil {
		return nil, err
	}

	productsList.Id = id

	return productsList, err
}
