package controllers

import (
	"github.com/minhaj/library-system/interfaces"
	"github.com/minhaj/library-system/models"
	"github.com/minhaj/library-system/services"
)

func getName(n interfaces.Namer) {
	println(n.GetName())
}

func Library() {
	john := models.NewUser(1, "John", "john@gmail.com")
	book := models.NewBook(1, "Go Lang", "Mark Henry")

	result := services.Borrowed(book, john)

	println(result)

}
