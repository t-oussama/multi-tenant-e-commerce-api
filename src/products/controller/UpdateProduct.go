package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func UpdateProduct(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	model := new(entities.Product)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	objId, idErr := primitive.ObjectIDFromHex(c.Param("id"))
	if idErr != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	var err error = service.UpdateProduct(objId, model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}
