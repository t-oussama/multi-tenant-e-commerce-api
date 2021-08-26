package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
)

func GetById(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	result, err := service.GetById(objId, config)

	if result == nil {
		c.JSON(http.StatusNotFound, struct{}{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, *result)
}
