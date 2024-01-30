package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/book/storage"
)

type BookServicer interface {
	Add(book models.BookDTO) error
	RentBook(userID, bookID int) error
	BackBook(userID, bookID int) error
	BookList() ([]models.Book, error)
}

type BookService struct {
	storage storage.BookStorager
}

func NewBookService(storage storage.BookStorager) *BookService {
	return &BookService{
		storage: storage,
	}
}

func (b *BookService) Add(book models.BookDTO) error {
	return b.storage.Add(book)
}

func (b *BookService) RentBook(userID, bookID int) error {
	return b.storage.RentBook(userID, bookID)
}

func (b *BookService) BackBook(userID int, bookID int) error {
	return b.storage.BackBook(userID, bookID)
}

func (b *BookService) BookList() ([]models.Book, error) {
	return b.storage.BookList()
}
