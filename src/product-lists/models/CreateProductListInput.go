package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateProductsListInput struct {
	Title    string               `bson:"title" json:"title"`
	Products []primitive.ObjectID `bson:"products" json:"products"`
}
