package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/product-lists/service"
	"store.tn/api/src/shared/constants"
)

func GetLists(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	result, err := service.GetLists(config)

	if result == nil {
		c.JSON(http.StatusOK, struct{}{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}
