package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"golang-web-starter/router"
)

func main() {
	router := router.Load()
	router.LoadHTMLGlob("templates/*")
	fmt.Println("test")
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	router.Run() // listen and serve on 0.0.0.0:8080
}