package service

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/author/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name   = "name"
	author = models.Author{
		FirstName: name,
	}
)

func TestAuthorService_Add(t *testing.T) {
	storageMock := mocks.NewAuthorStorager(t)
	storageMock.On("Add", author).Return(nil)

	authorService := NewAuthorService(storageMock)

	err := authorService.Add(author)

	assert.Nil(t, err)
}

func TestAuthorService_GetTop(t *testing.T) {
	storageMock := mocks.NewAuthorStorager(t)
	storageMock.On("GetTop").Return([]models.Author{author}, nil)

	authorService := NewAuthorService(storageMock)

	got, err := authorService.GetTop()

	assert.Nil(t, err)
	assert.NotNil(t, got)
}
