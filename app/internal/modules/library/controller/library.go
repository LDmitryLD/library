package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/errors"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/responder"
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/library/service"
)

type Libraryer interface {
	AddAuthor(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	RentBook(w http.ResponseWriter, r *http.Request)
	BackBook(w http.ResponseWriter, r *http.Request)
	GetTop(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	BookList(w http.ResponseWriter, r *http.Request)
}

type LibraryController struct {
	service service.LibraryServicer
	responder.Responder
}

func NewLibraryController(service service.LibraryServicer) *LibraryController {
	return &LibraryController{
		service:   service,
		Responder: &responder.Respond{},
	}
}

func (l *LibraryController) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var req AddAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq, err)
		l.ErrBadRequest(w, err)
		return
	}

	author := models.Author{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
	}

	if err := l.service.AddAuthor(author); err != nil {
		l.ErrInternal(w, err)
		return
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: "автор успешно добавлен",
	}

	l.OutputJSON(w, resp)
}

func (l *LibraryController) AddUser(w http.ResponseWriter, r *http.Request) {
	var req AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq, err)
		l.ErrBadRequest(w, err)
		return
	}

	user := models.UserDTO{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
	}

	if err := l.service.AddUser(user); err != nil {
		l.ErrInternal(w, err)
		return
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: "пользователь успешно добавлен",
	}

	l.OutputJSON(w, resp)
}

func (l *LibraryController) AddBook(w http.ResponseWriter, r *http.Request) {
	var req AddBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq, err)
		l.ErrBadRequest(w, err)
		return
	}

	book := models.BookDTO{
		AuthorID: req.AuthorID,
		Title:    req.Title,
	}

	if err := l.service.AddBook(book); err != nil {
		switch err {
		case errors.ErrAuthorNotFound:
			l.ErrNotFound(w, err)
			return
		default:
			l.ErrInternal(w, err)
			return
		}
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: "книга успешно добавлена",
	}

	l.OutputJSON(w, resp)
}

func (l *LibraryController) RentBook(w http.ResponseWriter, r *http.Request) {
	var req RentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq, err)
		l.ErrBadRequest(w, err)
		return
	}

	if err := l.service.RentBook(req.UserID, req.BookID); err != nil {
		switch err {
		case errors.ErrBookNotFound:
			l.ErrNotFound(w, err)
			return
		case errors.ErrUserNotFound:
			l.ErrNotFound(w, err)
			return
		default:
			l.ErrInternal(w, err)
			return
		}
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: "книга успешно взята в аренду",
	}

	l.OutputJSON(w, resp)
}

func (l *LibraryController) BackBook(w http.ResponseWriter, r *http.Request) {
	var req RentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq, err)
		l.ErrBadRequest(w, err)
		return
	}

	if err := l.service.BackBook(req.UserID, req.BookID); err != nil {
		switch err {
		case errors.ErrBookNotFound:
			l.ErrNotFound(w, err)
			return
		case errors.ErrUserNotFound:
			l.ErrNotFound(w, err)
			return
		case errors.ErrBookIsNotAvaliable:
			l.BookNotAvaliable(w, err)
			return
		default:
			l.ErrInternal(w, err)
			return
		}
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: "книга успешно возвращена",
	}

	l.OutputJSON(w, resp)
}

func (l *LibraryController) GetTop(w http.ResponseWriter, r *http.Request) {

	top, err := l.service.GetTop()
	if err != nil {
		l.ErrInternal(w, err)
		return
	}

	l.OutputJSON(w, top)
}

func (l *LibraryController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := l.service.GetUsersList()
	if err != nil {
		l.ErrInternal(w, err)
		return
	}

	l.OutputJSON(w, users)
}

func (l *LibraryController) BookList(w http.ResponseWriter, r *http.Request) {
	books, err := l.service.BookList()
	if err != nil {
		l.ErrInternal(w, err)
		return
	}

	l.OutputJSON(w, books)
}
