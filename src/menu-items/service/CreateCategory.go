package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/menu-items/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func CreateMenuItem(data *entities.Category, config *constants.TenantConfig) error {
	for i := 0; i < len(data.SubCategories); i++ {
		data.SubCategories[i].Id = primitive.NewObjectID()
		repository.InitializeSubCategoryProducts(data.SubCategories[i].Id, config)
	}

	id, err := repository.CreateCategory(*data, config)
	data.Id = id
	return err
}
