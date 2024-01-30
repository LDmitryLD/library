package storage

import (
	"projects/LDmitryLD/library/app/internal/db/adapter"
	"projects/LDmitryLD/library/app/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=BookStorager
type BookStorager interface {
	Add(book models.BookDTO) error
	RentBook(userID, bookID int) error
	BackBook(userID, bookID int) error
	BookList() ([]models.Book, error)
}

type BookStorage struct {
	adapter adapter.SQLAdapterer
}

func NewBookStorage(sqlAdapter *adapter.SQLAdapter) *BookStorage {
	return &BookStorage{
		adapter: sqlAdapter,
	}
}

func (s *BookStorage) Add(book models.BookDTO) error {
	return s.adapter.CreateBook(book)
}

func (s *BookStorage) RentBook(userID, bookID int) error {
	return s.adapter.TakeBook(userID, bookID)
}

func (s *BookStorage) BackBook(userID, bookID int) error {
	return s.adapter.BackBook(userID, bookID)
}

func (s *BookStorage) BookList() ([]models.Book, error) {
	return s.adapter.GetBookList()
}
