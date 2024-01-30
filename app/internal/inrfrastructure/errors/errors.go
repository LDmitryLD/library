package errors

import "fmt"

var (
	ErrAuthorNotFound     = fmt.Errorf("автор не найден")
	ErrUserNotFound       = fmt.Errorf("пользователь не найден")
	ErrBookNotFound       = fmt.Errorf("кннига не найдена")
	ErrBookIsNotAvaliable = fmt.Errorf("книга занята")
)
