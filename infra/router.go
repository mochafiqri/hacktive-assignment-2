package infra

import (
	"assignment2/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterInit() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.NoRoute(errorNotfound)
	r.NoMethod(errorNotfound)
	return r
}

func errorNotfound(c *gin.Context) {
	helper.JSON(c, http.StatusNotFound, "api not found", helper.Map{})
}
