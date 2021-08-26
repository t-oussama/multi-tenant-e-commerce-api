package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/products/models"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
)

func CreateProduct(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	model := new(models.CreateProductInput)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	var err error = service.CreateProduct(model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, model.Product)
}
