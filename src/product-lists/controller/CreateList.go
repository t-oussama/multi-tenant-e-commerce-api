package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/product-lists/models"
	"store.tn/api/src/product-lists/service"
	"store.tn/api/src/shared/constants"
)

func CreateList(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	model := new(models.CreateProductsListInput)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	result, err := service.CreateList(model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
