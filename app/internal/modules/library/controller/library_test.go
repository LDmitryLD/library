package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/errors"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/responder"
	"projects/LDmitryLD/library/app/internal/models"
	"projects/LDmitryLD/library/app/internal/modules/library/service/mocks"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var (
	title    = "title"
	name     = "test name"
	userID   = 1
	bookID   = 2
	authorID = 3
	errTest  = fmt.Errorf("test error")
	author   = models.Author{
		FirstName: name,
	}
	user = models.User{
		FirstName: name,
	}
	userDTO = models.UserDTO{
		FirstName: name,
	}
	book = models.Book{
		Title: title,
	}
	bookDTO = models.BookDTO{
		Title: title,
	}
	rentReq = RentRequest{
		UserID: userID,
		BookID: bookID,
	}
)

func TestLibraryController_AddAuthor_BadRequest(t *testing.T) {
	req := map[string]interface{}{"first_name": 1}
	reqJSON, _ := json.Marshal(req)

	libController := LibraryController{
		Responder: &responder.Respond{},
	}

	s := httptest.NewServer(http.HandlerFunc(libController.AddAuthor))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLibraryController_AddAuthor_Internal(t *testing.T) {
	req := AddAuthorRequest{
		FirstName: name,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddAuthor", author).Return(fmt.Errorf("test error"))

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddAuthor))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLibraryController_AddAuthor(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: "автор успешно добавлен",
	}
	req := AddAuthorRequest{
		FirstName: name,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddAuthor", author).Return(nil)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddAuthor))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var got models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expect, got)
}

func TestLibraryController_AddUser_BadRequest(t *testing.T) {
	req := map[string]interface{}{"first_name": 1}
	reqJSON, _ := json.Marshal(req)

	libController := LibraryController{
		Responder: &responder.Respond{},
	}

	s := httptest.NewServer(http.HandlerFunc(libController.AddUser))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLibraryController_AddUser_Internal(t *testing.T) {
	req := AddUserRequest{
		FirstName: name,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddUser", userDTO).Return(fmt.Errorf("test error"))

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddUser))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLibraryController_AddUser(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: "пользователь успешно добавлен",
	}
	req := AddUserRequest{
		FirstName: name,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddUser", userDTO).Return(nil)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddUser))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var got models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expect, got)
}

func TestLibraryController_AddBook_BadRequest(t *testing.T) {
	req := map[string]interface{}{"title": 1}
	reqJSON, _ := json.Marshal(req)

	libController := LibraryController{
		Responder: &responder.Respond{},
	}

	s := httptest.NewServer(http.HandlerFunc(libController.AddBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLibraryController_AddBook_Internal(t *testing.T) {
	req := AddBookRequest{
		Title: title,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddBook", bookDTO).Return(errTest)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLibraryController_AddBook(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: "книга успешно добавлена",
	}
	req := AddBookRequest{
		Title: title,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("AddBook", bookDTO).Return(nil)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.AddBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var got models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expect, got)
}

func TestLibraryController_RentBook_BadRequest(t *testing.T) {
	req := map[string]interface{}{"user_id": "1"}
	reqJSON, _ := json.Marshal(req)

	libController := LibraryController{
		Responder: &responder.Respond{},
	}

	s := httptest.NewServer(http.HandlerFunc(libController.RentBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLibraryController_RentBook_BookNotFound(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("RentBook", userID, bookID).Return(errors.ErrBookNotFound)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.RentBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLibraryController_RentBook_UserNotFound(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("RentBook", userID, bookID).Return(errors.ErrUserNotFound)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.RentBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLibraruController_RentBook_Internal(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("RentBook", userID, bookID).Return(errTest)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.RentBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLibraryController_RentBook(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: "книга успешно взята в аренду",
	}
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("RentBook", userID, bookID).Return(nil)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.RentBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var got models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, got)
}

func TestLibraryController_BackBook_BadRequest(t *testing.T) {
	req := map[string]interface{}{"user_id": "1"}
	reqJSON, _ := json.Marshal(req)

	libController := LibraryController{
		Responder: &responder.Respond{},
	}

	s := httptest.NewServer(http.HandlerFunc(libController.BackBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestLibraryController_BackBook_BookNotFound(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BackBook", userID, bookID).Return(errors.ErrBookNotFound)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.BackBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLibraryController_BackBook_UserNotFound(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BackBook", userID, bookID).Return(errors.ErrUserNotFound)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.BackBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLibraruController_BackBook_Internal(t *testing.T) {
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BackBook", userID, bookID).Return(errTest)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.BackBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestLibraryController_BackBook(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: "книга успешно возвращена",
	}
	reqJSON, _ := json.Marshal(rentReq)

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BackBook", userID, bookID).Return(nil)

	libController := NewLibraryController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(libController.BackBook))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var got models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, got)
}

func TestLibraryController_GetTop_Internal(t *testing.T) {
	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("GetTop").Return([]models.Author{}, errTest)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/authors/top", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/authors/top", libController.GetTop)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestLibraryController_GetTop(t *testing.T) {
	expect := []models.Author{author}

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("GetTop").Return([]models.Author{author}, nil)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/authors/top", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/authors/top", libController.GetTop)

	r.ServeHTTP(rr, req)

	var got []models.Author
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, got)
}

func TestLibraryController_GetUsers_Internal(t *testing.T) {
	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("GetUsersList").Return([]models.User{}, errTest)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/users", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/users", libController.GetUsers)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestLibraryController_GetUsers(t *testing.T) {
	expect := []models.User{user}

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("GetUsersList").Return([]models.User{user}, nil)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/users", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/users", libController.GetUsers)

	r.ServeHTTP(rr, req)

	var got []models.User
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, got)
}

func TestLibraryController_BookList_Internal(t *testing.T) {
	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BookList").Return([]models.Book{}, errTest)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/books", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/books", libController.BookList)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestLibraryController_BookList(t *testing.T) {
	expect := []models.Book{book}

	serviceMock := mocks.NewLibraryServicer(t)
	serviceMock.On("BookList").Return([]models.Book{book}, nil)

	libController := NewLibraryController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/library/books", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/library/books", libController.BookList)

	r.ServeHTTP(rr, req)

	var got []models.Book
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, got)
}
