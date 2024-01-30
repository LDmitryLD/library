package storage

import (
	"projects/LDmitryLD/library/app/internal/db/adapter"
	"projects/LDmitryLD/library/app/internal/models"
)

type UserStorager interface {
	Add(user models.UserDTO) error
	GetList() ([]models.User, error)
}

type UserStorage struct {
	adapter adapter.SQLAdapterer
}

func NewUserStorage(sqlAdapter *adapter.SQLAdapter) *UserStorage {
	return &UserStorage{
		adapter: sqlAdapter,
	}
}

func (u *UserStorage) Add(user models.UserDTO) error {
	return u.adapter.CreateUser(user)
}

func (u *UserStorage) GetList() ([]models.User, error) {
	return u.adapter.UserList()
}
