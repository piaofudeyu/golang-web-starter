package handler

import (
	"github.com/gin-gonic/gin"
	"golang-web-starter/model"
	"net/http"
	"strconv"
	"golang-web-starter/utility"
)

func ShowIndexPage(c *gin.Context)  {
	books := model.GetAllBooks()
	utility.Render(
		c,
		gin.H{
			"title": "Hi",
			"payload": books,
		},
		"index.html",
	)
}

func GetBook(c *gin.Context) {
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		if book, err := model.GetBookByID(bookID); err == nil {
			utility.Render(
				c,
				gin.H{
					"title": "Book",
					"payload": book,
				},
				"book.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}

}
