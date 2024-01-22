package errorlib

import (
	"errors"
	"net/http"
)

var ErrHttpInternal = errors.New("some internal error happened")
var ErrHttpBadRequest = errors.New("probably bad request")
var ErrHttpInvalidRequestData = errors.New("some data is invalid")

type JSONError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func GetJSONError(msg string, err error) *JSONError {
	var jsonErr = JSONError{
		Message: msg,
	}

	switch {
	case errors.Is(err, ErrHttpInternal):
		jsonErr.Code = http.StatusInternalServerError
	case errors.Is(err, ErrHttpBadRequest):
		jsonErr.Code = http.StatusBadRequest
	case errors.Is(err, ErrHttpInvalidRequestData):
		jsonErr.Code = http.StatusUnprocessableEntity
	default:
		jsonErr.Code = http.StatusInternalServerError
	}

	return &jsonErr
}
