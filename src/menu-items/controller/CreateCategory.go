package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/menu-items/service"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func CreateMenuItem(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	var model *entities.Category = new(entities.Category)
	if err := c.BindJSON(model); err != nil {
		return
	}

	var err error = service.CreateMenuItem(model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, model)
}
