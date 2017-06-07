package handler

import (
	"github.com/gin-gonic/gin"
	"golang-web-starter/model"
	"net/http"
	"strconv"
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

func GetBook(c *gin.Context) {
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		if book, err := model.GetBookByID(bookID); err == nil {
			c.HTML(
				http.StatusOK,
				"book.html",
				gin.H{
					"title": "Book",
					"book": book,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}

}
