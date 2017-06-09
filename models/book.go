package models

import (
	"golang-web-starter/utility"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" gorm:"type:varchar(100);unique"`
	Author string `json:"author" gorm:"type:varchar(100)"`
}

var bookList = []Book{
	{
		ID:     1,
		Title:  "Getting To YES",
		Author: "Roger Fisher",
	},
	{
		ID:     2,
		Title:  "You Don't Know JS",
		Author: "Getify",
	},
}

func GetAllBooks() []Book {
	var books []Book
	utility.DB().Find(&books)
	return books
}

func GetBookByID(id int) Book {
	var book Book
	utility.DB().First(&book, id)
	return book
}
