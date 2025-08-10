package models

type Profile struct {
	email string
}

type User struct {
	id   int
	name string
	Profile
	BorrowedBooks []Book
}

func NewUser(id int, name, email string) *User {
	return &User{
		id:      id,
		name:    name,
		Profile: Profile{email: email},
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmil() string {
	return u.Profile.email
}
