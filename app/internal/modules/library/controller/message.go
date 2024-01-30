package controller

const (
	LogErrDecodeReq = "ошибка при декодировании запроса: "
)

type AddAuthorRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type AddUserRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type AddBookRequest struct {
	AuthorID int    `json:"author_id"`
	Title    string `json:"title"`
}

type RentRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
