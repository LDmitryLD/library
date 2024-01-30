package controller

const (
	LogErrDecodeReq = "ошибка при декодировании запроса: "
)

type AddAuthorRequest struct {
	FirstName  string
	SecondName string
}

type AddUserRequest struct {
	FirstName  string
	SecondName string
}

type AddBookRequest struct {
	AuthorID int
	Title    string
}

type RentRequest struct {
	UserID int
	BookID int
}
