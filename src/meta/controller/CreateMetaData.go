package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/meta/service"
	"store.tn/api/src/shared/constants"
	"store.tn/api/src/shared/entities"
)

func CreateMetaData(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	model := new(entities.MetaData)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, struct{}{})
		return
	}

	err := service.CreateMetaData(model, config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, *model)
}
