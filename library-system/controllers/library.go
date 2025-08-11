package controllers

import (
	"fmt"

	"github.com/minhaj/library-system/services"
)

func Library() {
	books := services.ListOfBooks()
	users := services.ListOfUsers()

	for _, user := range users {
		fmt.Println(user.GetName())
	}

	for _, book := range books {
		fmt.Println(book.GetName())
	}

}
