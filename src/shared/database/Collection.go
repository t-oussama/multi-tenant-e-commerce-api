package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return Client.Database("storetn").Collection(name, opts...)
}
