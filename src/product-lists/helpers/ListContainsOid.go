package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/models"
)

func ListContainsOid(list []models.ProductBaseInfo, id primitive.ObjectID) bool {
	for _, el := range list {
		if el.Id == id {
			return true
		}
	}
	return false
}
