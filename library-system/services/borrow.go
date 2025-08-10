package services

import (
	"github.com/minhaj/library-system/models"
)

func Borrowed(book *models.Book, user *models.User) string {

	if book.BorrowedBy != nil {
		return "Book already borrowed!"
	}

	book.BorrowedBy = user

	return "Book borrowed by " + user.GetName()

}
