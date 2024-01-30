package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/user/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "name"
	userDTO = models.UserDTO{
		FirstName: name,
	}
)

func TestUserService_Add(t *testing.T) {
	storageMock := mocks.NewUserStorager(t)
	storageMock.On("Add", userDTO).Return(nil)

	userService := NewUserService(storageMock)

	err := userService.Add(userDTO)

	assert.Nil(t, err)
}

func TestUserService_GetList(t *testing.T) {
	storageMock := mocks.NewUserStorager(t)
	storageMock.On("GetList").Return([]models.User{{FirstName: name}}, nil)

	userService := NewUserService(storageMock)

	got, err := userService.GetList()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}
