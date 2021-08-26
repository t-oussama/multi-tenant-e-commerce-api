package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubCategoryProducts struct {
	Id       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Products []primitive.ObjectID `bson:"products"`
}
