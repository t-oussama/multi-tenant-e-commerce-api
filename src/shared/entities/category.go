package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label         string             `bson:"label" json:"label"`
	Icon          string             `bson:"icon" json:"icon"`
	SubCategories []SubCategory      `bson:"subCategories,omitempty" json:"subCategories,omitempty"`
}

type SubCategory struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label string             `bson:"label" json:"label"`
}
