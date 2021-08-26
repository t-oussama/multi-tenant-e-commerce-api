package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/products/service"
	"store.tn/api/src/shared/constants"
)

func GetAll(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	result, err := service.GetAll(config)

	if result == nil {
		c.JSON(http.StatusOK, struct{}{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}
