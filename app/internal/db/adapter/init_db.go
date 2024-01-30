package adapter

import (
	"log"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/helpers"
)

func (s *SQLAdapter) InitLibrary() {
	var count int
	if err := s.db.QueryRow("SELECT COUNT (*) FROM authors").Scan(&count); err != nil {
		log.Println("ошибка при проерке таблицы на пустоту ", err)
	}
	if count > 0 {
		return
	}

	authors := helpers.FakeAuthors()

	for _, author := range authors {
		if err := s.CreateAuthor(author); err != nil {
			log.Println("ошибка при создании автора: ", err)
			return
		}
	}

	users := helpers.FakeUsers()

	for _, user := range users {
		if err := s.CreateUser(user); err != nil {
			log.Println("ошибка при создании юзера: ", err)
		}
	}

	books := helpers.FakeBooks()

	for _, book := range books {
		if err := s.CreateBook(book); err != nil {
			log.Println("ошибка при создании книги: ", err)
		}
	}
}
