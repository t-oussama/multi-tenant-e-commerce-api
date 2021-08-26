package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label       string             `bson:"label" json:"label"`
	Images      []string           `bson:"images" json:"images"`
	Price       float64            `bson:"price" json:"price"`
	Reduction   float64            `bson:"reduction,omitempty" json:"reduction"`
	Description string             `bson:"description" json:"description"`
	IsAvailable bool               `bson:"isAvailable" json:"isAvailable"`
}
