package models

import "errors"

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
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
	return bookList
}

func GetBookByID(id int) (*book, error) {
	for _, book := range bookList {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}
