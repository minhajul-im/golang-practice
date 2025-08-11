package models

type Book struct {
	id         int
	name       string
	author     string
	BorrowedBy *User
}

func NewBook(id int, name, author string) *Book {
	return &Book{
		id:     id,
		name:   name,
		author: author,
	}
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetAuthor() string {
	return b.author
}
