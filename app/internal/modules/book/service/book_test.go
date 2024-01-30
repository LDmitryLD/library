package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/book/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	title   = "title"
	userID  = 1
	bookID  = 2
	bookDTO = models.BookDTO{
		Title: title,
	}
)

func TestBookService_Add(t *testing.T) {
	storageMock := mocks.NewBookStorager(t)
	storageMock.On("Add", bookDTO).Return(nil)

	bookService := NewBookService(storageMock)

	err := bookService.Add(bookDTO)

	assert.Nil(t, err)
}

func TestBookService_RentBook(t *testing.T) {
	storageMock := mocks.NewBookStorager(t)
	storageMock.On("RentBook", userID, bookID).Return(nil)

	bookService := NewBookService(storageMock)

	err := bookService.RentBook(userID, bookID)

	assert.Nil(t, err)
}

func TestBookService_BackBook(t *testing.T) {
	storageMock := mocks.NewBookStorager(t)
	storageMock.On("BackBook", userID, bookID).Return(nil)

	bookService := NewBookService(storageMock)

	err := bookService.BackBook(userID, bookID)

	assert.Nil(t, err)
}

func TestBookService_BookList(t *testing.T) {
	storageMock := mocks.NewBookStorager(t)
	storageMock.On("BookList").Return([]models.Book{{Title: title}}, nil)

	bookService := NewBookService(storageMock)

	got, err := bookService.BookList()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}
