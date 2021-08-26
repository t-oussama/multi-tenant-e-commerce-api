package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
)

func GetBySubCategoryId(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	subCategoryId, err := primitive.ObjectIDFromHex(c.Param("subCategoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	result, err := service.GetBySubCategoryId(subCategoryId, config)

	if result == nil {
		c.JSON(http.StatusOK, struct{}{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, *result)
}
