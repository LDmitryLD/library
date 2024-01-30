package models

type AuthorDTO struct {
	ID         int
	FirstName  string
	SecondName string
	RentCount  int
}

type Author struct {
	FirstName  string
	SecondName string
	RentCount  int
	Books      []BookForAuthor
}

type AuthorForBook struct {
	FirstName  string
	SecondName string
	RentCount  int
}
