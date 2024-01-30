package modules

import (
	aservice "projects/LDmitryLD/library/app/internal/modules/author/service"
	bookservice "projects/LDmitryLD/library/app/internal/modules/book/service"
	"projects/LDmitryLD/library/app/internal/modules/library/service"
	userservice "projects/LDmitryLD/library/app/internal/modules/user/service"
	"projects/LDmitryLD/library/app/internal/storages"
)

type Services struct {
	Library service.LibraryServicer
}

func NewServices(storages *storages.Storages) *Services {
	authorService := aservice.NewAuthorService(storages.Author)
	bookService := bookservice.NewBookService(storages.Book)
	userService := userservice.NewUserService(storages.User)

	libraryService := service.NewLibraryService(authorService, bookService, userService)

	return &Services{
		Library: libraryService,
	}
}
