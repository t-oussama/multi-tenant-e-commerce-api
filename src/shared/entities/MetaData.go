package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type MetaData struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
	Logo  string             `bson:"logo" json:"logo"`
}
