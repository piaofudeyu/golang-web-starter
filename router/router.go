package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Load loads the routes
func Load() *gin.Engine  {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Hi",
			},
		)
	})
	return router
}

