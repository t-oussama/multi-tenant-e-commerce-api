package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store.tn/api/src/product-lists/models"
	"store.tn/api/src/product-lists/service"
	"store.tn/api/src/shared/constants"
)

func UpdateList(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, new(interface{}))
		return
	}

	model := new(models.UpdateProductsListInput)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	err = service.UpdateList(id, model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, struct{}{})
}
