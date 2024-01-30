package helpers

import (
	"projects/LDmitryLD/library/app/internal/models"

	"github.com/brianvoe/gofakeit/v6"
)

func FakeAuthors() []models.Author {
	authors := make([]models.Author, 20)

	for i := 0; i < 20; i++ {
		author := models.Author{
			FirstName:  gofakeit.FirstName(),
			SecondName: gofakeit.LastName(),
			RentCount:  gofakeit.Number(1, 100),
		}
		authors[i] = author
	}

	return authors
}

func FakeUsers() []models.UserDTO {
	users := make([]models.UserDTO, 50)

	for i := 0; i < 50; i++ {
		user := models.UserDTO{
			FirstName:  gofakeit.FirstName(),
			SecondName: gofakeit.LastName(),
		}
		users[i] = user
	}

	return users
}

func FakeBooks() []models.BookDTO {
	books := make([]models.BookDTO, 100)

	for i := 0; i < 100; i++ {
		book := models.BookDTO{
			AuthorID: gofakeit.Number(1, 20),
			Title:    gofakeit.BookTitle(),
			Status:   models.StatusFree,
		}
		books[i] = book
	}

	return books
}
