package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
	"store.tn/api/src/shared/entities"
)

func GetStoreByPrefix(config *constants.TenantConfig, result *entities.MetaData) error {
	ctx := database.GetContext(30)
	return database.Collection(constants.CollectionNames.STORES).FindOne(ctx, bson.D{{"prefix", config.CollectionPrefix}}).Decode(result)
}
