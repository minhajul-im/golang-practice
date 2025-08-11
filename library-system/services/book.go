package services

import (
	"github.com/minhaj/library-system/models"
)

var (
	bookNames = []string{
		"Go Programming", "Clean Code", "Design Patterns",
		"The Go Workshop", "Concurrency in Go", "System Design",
		"Advanced Algorithms", "Refactoring", "Domain-Driven Design",
		"The Pragmatic Programmer",
	}

	bookAuthors = []string{
		"Alan Turing", "Robert Martin", "Erich Gamma",
		"Mark Henry", "Rob Pike", "Martin Fowler",
		"Donald Knuth", "Kent Beck", "Eric Evans",
		"Andy Hunt",
	}
)

func randomBookName() string {
	return bookNames[Random.Intn(len(bookNames))]
}

func randomBookAuthor() string {
	return bookAuthors[Random.Intn(len(bookAuthors))]
}

func ListOfBooks() []*models.Book {
	books := make([]*models.Book, 0, 20)

	for i := 0; i < 20; i++ {
		id := i
		authorName := randomBookAuthor()
		bookName := randomBookName()
		books = append(books, models.NewBook(id, authorName, bookName))
	}

	return books
}
