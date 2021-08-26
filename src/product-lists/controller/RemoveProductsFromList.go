package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/service"
	"store.tn/api/src/shared/constants"
)

func RemoveProductsFromList(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	model := make([]string, 0)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	oids := make([]primitive.ObjectID, len(model))
	for i, el := range model {
		oids[i], err = primitive.ObjectIDFromHex(el)
		if err != nil {
			c.JSON(http.StatusBadRequest, new(interface{}))
			return
		}
	}

	_, err = service.RemoveProductsFromList(objId, oids, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, struct{}{})
}
