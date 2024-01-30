package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/author/storage"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=AuthorServicer
type AuthorServicer interface {
	Add(author models.Author) error
	GetTop() ([]models.Author, error)
}

type AuthorService struct {
	storage storage.AuthorStorager
}

func NewAuthorService(storage storage.AuthorStorager) *AuthorService {
	return &AuthorService{
		storage: storage,
	}
}

func (a *AuthorService) Add(author models.Author) error {
	return a.storage.Add(author)
}

func (a *AuthorService) GetTop() ([]models.Author, error) {
	return a.storage.GetTop()
}
