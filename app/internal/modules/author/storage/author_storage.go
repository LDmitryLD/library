package storage

import (
	"projects/LDmitryLD/library/app/internal/db/adapter"
	"projects/LDmitryLD/library/app/internal/models"
)

type AuthorStorager interface {
	Add(author models.Author) error
	GetTop() ([]models.Author, error)
}

func NewAuthorStorage(sqlAdapter *adapter.SQLAdapter) *AuthorStorage {
	return &AuthorStorage{
		adapter: sqlAdapter,
	}
}

type AuthorStorage struct {
	adapter adapter.SQLAdapterer
}

func (a *AuthorStorage) Add(author models.Author) error {
	return a.adapter.CreateAuthor(author)
}

func (a *AuthorStorage) GetTop() ([]models.Author, error) {
	return a.adapter.GetTop()
}
