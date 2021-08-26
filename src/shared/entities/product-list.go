package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/models"
)

type ProductList struct {
	Id       primitive.ObjectID       `bson:"_id" json:"id"`
	Title    string                   `bson:"title" json:"title"`
	Products []models.ProductBaseInfo `bson:"products" json:"products"`
}
