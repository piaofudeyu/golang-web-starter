package models

import (
	"errors"
	"golang-web-starter/utility"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" gorm:"type:varchar(100);unique"`
	Author string `json:"author" gorm:"type:varchar(100)"`
}

var bookList = []book{
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

func GetAllBooks() []book {
	var books []book
	utility.DB().Find(&books)
	return books
}

func GetBookByID(id int) (*book, error) {
	for _, book := range bookList {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}
