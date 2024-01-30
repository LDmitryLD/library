package models

import "time"

type RentedBooksDTO struct {
	UserID     int
	BookID     int
	BorrowDate time.Time
	ReturnDate time.Time
}
