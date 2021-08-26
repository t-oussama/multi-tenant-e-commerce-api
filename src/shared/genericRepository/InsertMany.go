package genericRepository

import (
	"fmt"

	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/database"
)

func InsertMany(collection string, data []interface{}, config *constants.TenantConfig) {
	ctx := database.GetContext(30)
	res, insertErr := database.Collection(config.CollectionPrefix+collection).InsertMany(ctx, data)
	if insertErr != nil {
		panic(insertErr)
	}
	fmt.Println(res)
}
