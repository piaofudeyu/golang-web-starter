package model

import "testing"

func TestGetAllBooks(t *testing.T) {
	books := GetAllBooks()

	if len(books) != len(bookList) {
		t.Fail()
	}

	for i, v := range books {
		if v.Title != bookList[i].Title ||
			v.ID != bookList[i].ID ||
			v.Author != bookList[i].Title {

			t.Fail()
			break
		}
	}
}