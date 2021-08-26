package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/service"
	"store.tn/api/src/shared/constants"
)

func DeleteList(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	err = service.DeleteList(objId, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, struct{}{})
}
