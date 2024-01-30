package models

type UserDTO struct {
	ID         int
	FirstName  string
	SecondName string
}

type User struct {
	FirstName   string
	SecondName  string
	RentedBooks []Book
}
