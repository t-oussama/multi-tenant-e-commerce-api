package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductBaseInfo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label     string             `bson:"label" json:"label"`
	Image     string             `bson:"images" json:"images"`
	Price     float64            `bson:"price" json:"price"`
	Reduction float64            `bson:"reduction,omitempty" json:"reduction"`
}
