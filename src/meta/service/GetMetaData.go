package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"store.tn/api/src/meta/repository"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func GetMetaData(config *constants.TenantConfig) (*entities.MetaData, error) {
	var result *entities.MetaData = new(entities.MetaData)

	err := repository.GetStoreByPrefix(config, result)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return result, err
}
