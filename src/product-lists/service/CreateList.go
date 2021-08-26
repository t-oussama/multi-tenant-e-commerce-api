package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/models"
	"store.tn/api/src/product-lists/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
	sharedModels "store.tn/api/src/shared/models"
)

func CreateList(data *models.CreateProductsListInput, config *constants.TenantConfig) (*entities.ProductList, error) {

	productsList := new(entities.ProductList)

	productsList.Id = primitive.NewObjectID()
	productsList.Title = data.Title
	productsList.Products = make([]sharedModels.ProductBaseInfo, 0)

	findError := repository.GetProductsByIdsList(data.Products, &productsList.Products, config)

	if findError != nil {
		return nil, findError
	}

	id, err := genericRepository.InsertOne(constants.CollectionNames.PRODUCT_LISTS, *productsList, config)

	if err != nil {
		return nil, err
	}

	productsList.Id = id

	return productsList, err
}
