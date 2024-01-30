package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	aservice "projects/LDmitryLD/library/app/internal/modules/author/service"
	bookservice "projects/LDmitryLD/library/app/internal/modules/book/service"
	userservice "projects/LDmitryLD/library/app/internal/modules/user/service"
)

type LibraryServicer interface {
	AddAuthor(author models.Author) error
	AddUser(user models.UserDTO) error
	AddBook(book models.BookDTO) error
	RentBook(userID, bookID int) error
	BackBook(userID, bookID int) error
	GetTop() ([]models.Author, error)
	GetUsersList() ([]models.User, error)
	BookList() ([]models.Book, error)
}

type LibraryService struct {
	Author aservice.AuthorServicer
	Book   bookservice.BookServicer
	User   userservice.UserServicer
}

func NewLibraryService(author aservice.AuthorServicer, book bookservice.BookServicer, user userservice.UserServicer) *LibraryService {
	return &LibraryService{
		Author: author,
		Book:   book,
		User:   user,
	}
}

func (l *LibraryService) AddAuthor(author models.Author) error {
	return l.Author.Add(author)
}

func (l *LibraryService) AddUser(user models.UserDTO) error {
	return l.User.Add(user)
}

func (l *LibraryService) AddBook(book models.BookDTO) error {
	return l.Book.Add(book)
}

func (l *LibraryService) RentBook(userID, bookID int) error {
	return l.Book.RentBook(userID, bookID)
}

func (l *LibraryService) BackBook(userID, bookID int) error {
	return l.Book.BackBook(userID, bookID)
}

func (l *LibraryService) GetTop() ([]models.Author, error) {
	return l.Author.GetTop()
}

func (l *LibraryService) GetUsersList() ([]models.User, error) {
	return l.User.GetList()
}

func (l *LibraryService) BookList() ([]models.Book, error) {
	return l.Book.BookList()
}
