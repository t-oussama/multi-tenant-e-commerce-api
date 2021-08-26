package service

import (
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
	"store.tn/api/src/shared/genericRepository"
)

func CreateMetaData(data *entities.MetaData, config *constants.TenantConfig) error {

	id, err := genericRepository.InsertOne(constants.CollectionNames.META_DATA, data, config)

	if err != nil {
		return err
	}

	data.Id = id

	return nil
}
