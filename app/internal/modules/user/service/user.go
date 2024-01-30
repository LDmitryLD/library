package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/user/storage"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=UserServicer
type UserServicer interface {
	Add(user models.UserDTO) error
	GetList() ([]models.User, error)
}

type UserService struct {
	storage storage.UserStorager
}

func NewUserService(storage storage.UserStorager) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u *UserService) Add(user models.UserDTO) error {
	return u.storage.Add(user)
}

func (u *UserService) GetList() ([]models.User, error) {
	return u.storage.GetList()
}
