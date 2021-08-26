package genericRepository

import (
	"go.mongodb.org/mongo-driver/bson"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func Find(collection string, query bson.D, results interface{}, config *constants.TenantConfig) error {
	ctx := database.GetContext(30)
	cur, err := database.Collection(config.CollectionPrefix+collection).Find(ctx, query)
	if err != nil {
		cur.Close(ctx)
		return err
	}

	err = cur.All(ctx, results)
	cur.Close(ctx)
	return err
}
