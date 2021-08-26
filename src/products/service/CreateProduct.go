package service

import (
	"store.tn/api/src/products/models"
	"store.tn/api/src/products/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/genericRepository"
)

func CreateProduct(data *models.CreateProductInput, config *constants.TenantConfig) error {
	product := &data.Product
	subCategories := data.SubCategories

	id, err := genericRepository.InsertOne(constants.CollectionNames.PRODUCTS, *product, config)

	if err != nil {
		return err
	}
	product.Id = id

	for i := 0; i < len(subCategories); i++ {
		err = repository.AddProductToSubCategory(product.Id, subCategories[i], config)
	}

	return err
}
