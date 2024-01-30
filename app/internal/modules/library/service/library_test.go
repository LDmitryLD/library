package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	amocks "projects/LDmitryLD/library/app/internal/modules/author/service/mocks"
	bmocks "projects/LDmitryLD/library/app/internal/modules/book/service/mocks"
	umocks "projects/LDmitryLD/library/app/internal/modules/user/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name   = "name"
	title  = "title"
	bookID = 1
	userID = 2
	author = models.Author{
		FirstName: name,
	}
	user = models.User{
		FirstName: name,
	}
	userDTO = models.UserDTO{
		FirstName: name,
	}
	book = models.Book{
		Title: title,
	}
	bookDTO = models.BookDTO{
		Title: title,
	}
)

func TestLibraryService_AddAuthor(t *testing.T) {
	authorService := amocks.NewAuthorServicer(t)
	bookService := bmocks.NewBookServicer(t)
	userService := umocks.NewUserServicer(t)

	authorService.On("Add", author).Return(nil)

	libService := NewLibraryService(authorService, bookService, userService)

	err := libService.AddAuthor(author)

	assert.Nil(t, err)
}

func TestLibraryService_AddUser(t *testing.T) {
	userService := umocks.NewUserServicer(t)

	userService.On("Add", userDTO).Return(nil)

	libService := LibraryService{
		User: userService,
	}
	err := libService.AddUser(userDTO)

	assert.Nil(t, err)
}

func TestLibraryService_AddBook(t *testing.T) {
	bookService := bmocks.NewBookServicer(t)

	bookService.On("Add", bookDTO).Return(nil)

	libService := LibraryService{
		Book: bookService,
	}
	err := libService.AddBook(bookDTO)

	assert.Nil(t, err)
}

func TestLibraryService_RentBook(t *testing.T) {
	bookService := bmocks.NewBookServicer(t)

	bookService.On("RentBook", userID, bookID).Return(nil)

	libService := LibraryService{
		Book: bookService,
	}
	err := libService.RentBook(userID, bookID)

	assert.Nil(t, err)
}

func TestLibraryService_BackBook(t *testing.T) {
	bookService := bmocks.NewBookServicer(t)

	bookService.On("BackBook", userID, bookID).Return(nil)

	libService := LibraryService{
		Book: bookService,
	}
	err := libService.BackBook(userID, bookID)

	assert.Nil(t, err)
}

func TestLibraryService_GetTop(t *testing.T) {
	authorService := amocks.NewAuthorServicer(t)

	authorService.On("GetTop").Return([]models.Author{author}, nil)

	libService := LibraryService{
		Author: authorService,
	}

	got, err := libService.GetTop()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestLibraryService_GetUsersList(t *testing.T) {
	userService := umocks.NewUserServicer(t)

	userService.On("GetList").Return([]models.User{user}, nil)

	libService := LibraryService{
		User: userService,
	}

	got, err := libService.GetUsersList()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}

func TestLibraryService_BookList(t *testing.T) {
	bookService := bmocks.NewBookServicer(t)

	bookService.On("BookList").Return([]models.Book{book}, nil)

	libService := LibraryService{
		Book: bookService,
	}

	got, err := libService.BookList()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}
