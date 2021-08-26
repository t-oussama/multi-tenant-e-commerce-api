package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/shared/entities"
)

type CreateProductInput struct {
	Product       entities.Product     `json:"product"`
	SubCategories []primitive.ObjectID `json:"subCategories"`
}
