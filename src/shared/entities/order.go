package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Address        string               `bson:"address" json:"address"`
	UserId         primitive.ObjectID   `bson:"userId" json:"userId"`
	Porducts       []primitive.ObjectID `bson:"products" json:"products"`
	ProductsPrices []OrderProductPrice  `bson:"productsPrices" json:"productsPrices"`
	State          int                  `bson:"state" json:"state"` // 0: pending, 1 delivered, 2 cancelled
}

type OrderProductPrice struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Price float64            `bson:"price" json:"price"`
}
