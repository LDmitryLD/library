package adapter

import (
	"log"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/errors"
	"projects/LDmitryLD/library/app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type SQLAdapterer interface {
	CreateAuthor(author models.Author) error
	CreateBook(book models.BookDTO) error
	CreateUser(user models.UserDTO) error
	TakeBook(userID, bookID int) error
	BackBook(userID, bookID int) error
	GetTop() ([]models.Author, error)
	UserList() ([]models.User, error)
	GetBookList() ([]models.Book, error)
}

type SQLAdapter struct {
	db *sqlx.DB
}

func NewSQLAdapter(db *sqlx.DB) *SQLAdapter {
	return &SQLAdapter{
		db: db,
	}
}

func (s *SQLAdapter) CreateAuthor(author models.Author) error {
	q := `
	INSERT INTO authors 
		(first_name, second_name, rent_count)
	VALUES
		($1, $2, $3)
	RETURNING id		
	`
	var id int
	if err := s.db.QueryRow(q, author.FirstName, author.SecondName, author.RentCount).Scan(&id); err != nil {
		log.Println("ошибка при записи автора в бд: ", err)
		return err
	}

	log.Printf("автор с id %d записан в БД\n", id)
	return nil
}

func (s *SQLAdapter) CreateBook(book models.BookDTO) error {
	if _, err := s.getAuthorByID(book.AuthorID); err != nil {
		return err
	}

	q := `
	INSERT INTO books
		(author_id, title, status)
	VALUES
		($1, $2, $3)	
	RETURNING id	
	`
	var id int
	if err := s.db.QueryRow(q, book.AuthorID, book.Title, models.StatusFree).Scan(&id); err != nil {
		log.Println("ошибка при записи книги в бд: ", err)
		return err
	}

	log.Printf("книга с id %d записана в бд\n", id)

	return nil
}

func (s *SQLAdapter) CreateUser(user models.UserDTO) error {
	q := `
	INSERT INTO users
		(first_name, second_name)
	VALUES
		($1, $2)
	RETURNING id		
	`
	var id int
	if err := s.db.QueryRow(q, user.FirstName, user.SecondName).Scan(&id); err != nil {
		log.Println("ошибка при записи пользователя в бд:", err)
		return err
	}

	log.Printf("Пользователь с id %d записан в бд\n", id)

	return nil
}

func (s *SQLAdapter) TakeBook(userID, bookID int) error {
	if err := s.getUserByID(userID); err != nil {
		return err
	}

	book, err := s.getBookByID(bookID)
	if err != nil {
		return err
	}

	if !book.IsAvaliable() {
		return errors.ErrBookIsNotAvaliable
	}

	q := `
	INSERT INTO rented_books
		(user_id, book_id, borrow_date)
	VALUES
		($1, $2, $3)	
	`
	_, err = s.db.Exec(q, userID, bookID, time.Now())
	if err != nil {
		log.Println("ошибка при получении книги в аренду: ", err)
		return err
	}

	if err := s.changeBookStatus(bookID, models.StatusInRent); err != nil {
		return err
	}

	s.incRentCount(book.AuthorID)

	return nil
}

func (s *SQLAdapter) BackBook(userID, bookID int) error {
	if err := s.getUserByID(userID); err != nil {
		return err
	}

	err := s.checkBook(bookID)
	if err != nil && err != errors.ErrBookIsNotAvaliable {
		return err
	}

	q := `
	DELETE FROM rented_books
	WHERE book_id = ($1)
	`

	_, err = s.db.Exec(q, bookID)
	if err != nil {
		log.Println("ошибка при сдаче книги: ", err)
		return err
	}

	if err := s.changeBookStatus(bookID, models.StatusFree); err != nil {
		return err
	}

	return nil
}

func (s *SQLAdapter) GetBookList() ([]models.Book, error) {
	q := `
	SELECT * FROM books
	`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	var booksDTO []models.BookDTO
	for rows.Next() {
		var b models.BookDTO
		err := rows.Scan(&b.ID, &b.AuthorID, &b.Title, &b.Status)
		if err != nil {
			return nil, err
		}

		booksDTO = append(booksDTO, b)
	}

	books := make([]models.Book, len(booksDTO))
	for i, book := range booksDTO {
		books[i].Title = book.Title

		author, err := s.getAuthorByID(book.AuthorID)
		if err != nil {
			return nil, err
		}
		books[i].Author.FirstName = author.FirstName
		books[i].Author.SecondName = author.SecondName
		books[i].Author.RentCount = author.RentCount
	}

	return books, nil
}

func (s *SQLAdapter) GetTop() ([]models.Author, error) {
	q := `
	SELECT 
		id, first_name, second_name, rent_count 
	FROM 
		authors
	ORDER BY 
		rent_count
	DESC LIMIT 10
			
	`

	rows, err := s.db.Query(q)
	if err != nil {
		log.Println("ошибка при получении сипска авторов ", err)
		return nil, err
	}
	defer rows.Close()

	authors := make([]models.Author, 10)
	var i int
	for rows.Next() {
		var id int
		var author models.Author
		err = rows.Scan(&id, &author.FirstName, &author.SecondName, &author.RentCount)
		if err != nil {
			log.Println("ошибка при получении сипска авторов ", err)
			return nil, err
		}

		var books []models.BookDTO
		if err := s.db.Select(&books, `SELECT * FROM books WHERE author_id = $1`, id); err != nil {
			return nil, err
		}

		authors[i] = author
		for _, book := range books {
			authors[i].Books = append(authors[i].Books, models.BookForAuthor{Title: book.Title})
		}
		i++
	}

	return authors, nil
}

func (s *SQLAdapter) UserList() ([]models.User, error) {
	q := ` 
	SELECT * FROM users
	`
	rows, err := s.db.Query(q)
	if err != nil {
		log.Println("ошибка при получении всех пользователей ", err)
		return nil, err
	}

	var usersDTO []models.UserDTO
	for rows.Next() {
		var u models.UserDTO
		err := rows.Scan(&u.ID, &u.FirstName, &u.SecondName)
		if err != nil {
			log.Println("ошибка при сканировании результатов")
		}
		usersDTO = append(usersDTO, u)
	}

	users := make([]models.User, len(usersDTO))
	for i, userDTO := range usersDTO {
		users[i].FirstName = userDTO.FirstName
		users[i].SecondName = userDTO.SecondName
		books, err := s.getBookByUserID(userDTO.ID)
		if err != nil {
			return nil, err
		}
		for _, book := range books {

			author, err := s.getAuthorByID(book.AuthorID)
			if err != nil {
				return nil, err
			}
			users[i].RentedBooks = append(users[i].RentedBooks,
				models.Book{Title: book.Title, Author: models.AuthorForBook{FirstName: author.FirstName, SecondName: author.SecondName, RentCount: author.RentCount}})
		}
	}

	return users, nil
}

func (s *SQLAdapter) getBookByUserID(userID int) ([]models.BookDTO, error) {
	q := `
	SELECT 
		b.id, b.author_id, b.title
	FROM 
		rented_books rb	
	JOIN 
		books b ON rb.book_id = b.id
	WHERE rb.user_id = $1		
	`

	rows, err := s.db.Query(q, userID)
	if err != nil {
		return nil, err
	}

	var books []models.BookDTO
	for rows.Next() {
		var b models.BookDTO
		err := rows.Scan(&b.ID, &b.AuthorID, &b.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (s *SQLAdapter) getUserByID(userID int) error {
	q := `
	SELECT id FROM users WHERE id = $1
	`
	row := s.db.QueryRow(q, userID)

	var id int
	if err := row.Scan(&id); err != nil {
		log.Println("пользователь не найден", err)
		return errors.ErrUserNotFound
	}

	return nil
}

func (s *SQLAdapter) getBookByID(bookID int) (models.BookDTO, error) {
	q := `
	SELECT * FROM books WHERE id = $1
	`
	row := s.db.QueryRow(q, bookID)

	var book models.BookDTO
	if err := row.Scan(&book.ID, &book.AuthorID, &book.Title, &book.Status); err != nil {
		log.Println("кнгиа не найден", err)
		return models.BookDTO{}, errors.ErrBookNotFound
	}

	return book, nil
}

func (s *SQLAdapter) changeBookStatus(bookID int, status string) error {
	q := `
	UPDATE books
	SET status = $1
	WHERE id = $2
	`
	_, err := s.db.Exec(q, status, bookID)
	if err != nil {
		log.Println("ошибка при смене статуса книги: ", err)
		return err
	}

	log.Println("статус книги изменён")

	return nil
}

func (s *SQLAdapter) checkBook(bookID int) error {
	book, err := s.getBookByID(bookID)
	if err != nil {
		return err
	}

	if !book.IsAvaliable() {
		return errors.ErrBookIsNotAvaliable
	}

	return nil
}

func (s *SQLAdapter) getAuthorByID(authorID int) (models.Author, error) {
	q := `
	SELECT first_name, second_name, rent_count FROM authors WHERE id = $1
	`
	row := s.db.QueryRow(q, authorID)

	var author models.Author
	if err := row.Scan(&author.FirstName, &author.SecondName, &author.RentCount); err != nil {
		log.Println("автор не найден", err)
		return models.Author{}, errors.ErrAuthorNotFound
	}

	return author, nil
}

func (s *SQLAdapter) incRentCount(authotID int) {
	q := `
	UPDATE authors
	SET rent_count=rent_count + 1
	WHERE id = $1
	`

	_, err := s.db.Exec(q, authotID)
	if err != nil {
		log.Println("не получилось увеличить rent_count: ", err)
		return
	}

	log.Println("рейтинг автора увеличен успешно")
}
