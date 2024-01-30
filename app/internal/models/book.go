package models

const (
	StatusFree   = "free"
	StatusInRent = "in rent"
)

type BookDTO struct {
	ID       int
	AuthorID int `db:"author_id"`
	Title    string
	Status   string
}

func (b BookDTO) IsAvaliable() bool {
	return b.Status == StatusFree
}

type Book struct {
	Title  string
	Author AuthorForBook
}

type BookForAuthor struct {
	Title string
}
