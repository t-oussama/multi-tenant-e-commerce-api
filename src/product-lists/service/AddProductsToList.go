package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/helpers"
	"store.tn/api/src/product-lists/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
	"store.tn/api/src/shared/models"
)

func AddProductsToList(id primitive.ObjectID, ids []primitive.ObjectID, config *constants.TenantConfig) (*entities.ProductList, error) {

	productsList := new(entities.ProductList)

	findError := genericRepository.FindById(constants.CollectionNames.PRODUCT_LISTS, id, productsList, config)

	if findError != nil {
		return nil, findError
	}

	idsToAdd := make([]primitive.ObjectID, 0)

	for i := 0; i < len(ids); i++ {
		if helpers.ListContainsOid(productsList.Products, ids[i]) {
			idsToAdd = append(idsToAdd, ids[i])
		}
	}

	products := make([]models.ProductBaseInfo, 0)

	findError = repository.GetProductsByIdsList(idsToAdd, &products, config)

	if findError != nil {
		return nil, findError
	}

	productsList.Products = append(productsList.Products, products...)

	id, err := genericRepository.UpdateByID(constants.CollectionNames.PRODUCT_LISTS, id, *productsList, config)

	if err != nil {
		return nil, err
	}

	productsList.Id = id

	return productsList, err
}
