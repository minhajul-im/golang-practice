package models

type Book struct {
	id         int
	Title      string
	Author     string
	BorrowedBy *User
}

func NewBook(id int, title, author string) *Book {
	return &Book{
		id:     id,
		Title:  title,
		Author: author,
	}
}

func (b *Book) GetName() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}
