package errors

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewError(err error) *RestError {
	if err == sql.ErrNoRows {
		return NewNotFundError(err.Error())
	} else {
		sqlErr, ok := err.(*pq.Error)
		if !ok {
			return NewInternalServerError(err.Error())
		}

		return NewInternalServerError(fmt.Sprintf("Fatal: %t, Code: %s, Detail: %s", sqlErr.Fatal(), sqlErr.Code, sqlErr.Detail))
	}

}
