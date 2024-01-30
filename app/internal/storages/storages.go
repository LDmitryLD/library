package storages

import (
	"projects/LDmitryLD/library/app/internal/db/adapter"
	astorage "projects/LDmitryLD/library/app/internal/modules/author/storage"
	bookstorage "projects/LDmitryLD/library/app/internal/modules/book/storage"
	userstorage "projects/LDmitryLD/library/app/internal/modules/user/storage"
)

type Storages struct {
	Author astorage.AuthorStorager
	Book   bookstorage.BookStorager
	User   userstorage.UserStorager
}

func NewStorages(sqlAdapter *adapter.SQLAdapter) *Storages {
	return &Storages{
		Author: astorage.NewAuthorStorage(sqlAdapter),
		Book:   bookstorage.NewBookStorage(sqlAdapter),
		User:   userstorage.NewUserStorage(sqlAdapter),
	}
}
