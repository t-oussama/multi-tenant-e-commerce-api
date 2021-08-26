package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
)

func AddProductToSubCategory(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	productId, productIdErr := primitive.ObjectIDFromHex(c.Param("productId"))
	if productIdErr != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	subCategoryId, subCategoryIdErr := primitive.ObjectIDFromHex(c.Param("subCategoryId"))
	if subCategoryIdErr != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	err := service.AddProductToSubCategory(productId, subCategoryId, config)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, struct{}{})
		log.Fatal(err)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}
