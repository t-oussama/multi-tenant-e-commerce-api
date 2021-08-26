package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/entities"
)

type PopulatedSubCategoryProducts struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Products []entities.Product `bson:"products"`
}
