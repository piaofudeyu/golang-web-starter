package handler

import (
	"github.com/gin-gonic/gin"
	"golang-web-starter/model"
	"net/http"
)

func ShowIndexPage(c *gin.Context)  {
	books := model.GetAllBooks()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Hi",
			"books": books,
		},
	)
}
