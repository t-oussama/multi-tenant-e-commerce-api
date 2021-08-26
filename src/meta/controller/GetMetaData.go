package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"store.tn/api/src/meta/service"
	"store.tn/api/src/shared/constants"
)

func GetMetaData(c *gin.Context) {
	config := constants.TenantsConfig[c.Request.Host]

	result, err := service.GetMetaData(config)

	if result == nil {
		c.JSON(http.StatusNotFound, struct{}{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, struct{}{})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, *result)
}
