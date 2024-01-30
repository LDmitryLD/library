package docs

import (
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/library/controller"
)

// swagger:route POST /library/book user AddUserRequest
// Добавление пользователя.
// responses:
// 	200: AddUserResponse

// swagger:parameters AddUserRequest
type AddUserRequest struct {
	// required:true
	// in:body
	Body controller.AddUserRequest
}

// swagger:response AddUserResponse
type AddUserResponse struct {
	// in:body
	Body models.ApiResponse
}

// swagger:route GET /library/users user
// Получения списка пользователей.
// responses:
//  200: GetUsersResponse

// swagger:response GetUsersResponse
type GetUsersResponse struct {
	// in:body
	Body []models.User
}

// swagger:route POST /library/author author AddAuthorRequest
// Добавление автора.
// responses:
//  200: AddAuthorResponse

// swagger:parameters AddAuthorRequest
type AddAuthorRequest struct {
	// required:true
	// in:body
	Body controller.AddAuthorRequest
}

// swagger:response AddAuthorResponse
type AddAuthorResponse struct {
	// in:body
	Body models.ApiResponse
}

// swagger:route GET /library/authors/top author
// Топ 10 читаемых авторов.
// responses:
//  200: GetTopResponse

// swagger:response GetTopResponse
type GetTopResponse struct {
	// in:body
	Body []models.Author
}

// swagger:route POST /library/book book AddBookRequest
// Добавить книгу.
// responses:
//  200: AddBookResponse

// swagger:parameters AddBookRequest
type AddBookRequest struct {
	// required:true
	// in:body
	Body controller.AddBookRequest
}

// swagger:response AddBookResponse
type AddBookResponse struct {
	// in:body
	Body models.ApiResponse
}

// swagger:route PUT /library/book/rent book RentBookRequest
// Взять книгу.
// responses:
//  200: RentBookResponse

// swagger:parameters RentBookRequest
type RentBookRequest struct {
	// required:true
	// in:body
	Body controller.RentRequest
}

// swagger:response RentBookResponse
type RentBookResponse struct {
	// in:body
	Body models.Book
}

// swagger:route PUT /library/book/back book BackBookRequest
// Вернуть книгу.
// responses:
//  200: BackBookResponse

// swagger:parameters BackBookRequest
type BackBookRequest struct {
	// required:true
	// in:body
	Body controller.RentRequest
}

// swagger:response BackBookResponse
type BackBookResponse struct {
	// in:body
	Body models.ApiResponse
}
