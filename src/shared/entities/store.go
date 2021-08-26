package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Logo   string             `bson:"logo" json:"logo"`
	Prefix string             `bson:"prefix" json:"prefix"`
	Shops  []Shop             `bson:"shops" json:"shops"`
}

type Shop struct {
	Address     string `bson:"address" json:"address"`
	Phone       string `bson:"phone" json:"phone"`
	GeoLocation string `bson:"geoLocation" json:"geoLocation"`
}
