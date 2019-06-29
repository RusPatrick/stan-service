package errors

import (
	"net/http"

	"github.com/ruspatrick/stan-svc/infrastructure/helpers"
)

type Error struct {
	Err        error  `json:"-"`
	Type       string `json:"type,omitempty"`
	Source     string `json:"source,omitempty"`
	Title      string `json:"title,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

const (
	ServerType   = "server"
	BusinessType = "business"
)

func (e Error) Error() string {
	return e.Message
}

func CreateServerError(err error, title, message string) Error {
	return Error{
		Err:        err,
		Title:      title,
		Message:    message,
		Source:     helpers.GetFunctionName(2),
		Type:       ServerType,
		StatusCode: http.StatusInternalServerError,
	}
}

func CreateBusinessError(err error, title, message string) Error {
	return Error{
		Err:        err,
		Title:      title,
		Message:    message,
		Source:     helpers.GetFunctionName(2),
		Type:       BusinessType,
		StatusCode: http.StatusBadRequest,
	}
}
