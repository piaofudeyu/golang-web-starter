package router

import (
	"github.com/gin-gonic/gin"
	"golang-web-starter/handler"
)

// Load loads the routes
func Load() *gin.Engine  {
	router := gin.Default()

	router.GET("/", handler.ShowIndexPage)
	return router
}

