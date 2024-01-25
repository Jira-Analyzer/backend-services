package errorlib

import (
	"errors"
	"net/http"
)

var ErrHttpInternal = errors.New("some internal error happened")
var ErrHttpBadRequest = errors.New("probably bad request")
var ErrHttpInvalidRequestData = errors.New("some data is invalid")

type JSONError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	} `json:"error"`
}

func GetJSONError(msg string, err error) *JSONError {
	var jsonErr = JSONError{}
	jsonErr.Error.Message = msg
	switch {
	case errors.Is(err, ErrHttpInternal):
		jsonErr.Error.Code = http.StatusInternalServerError
	case errors.Is(err, ErrHttpBadRequest):
		jsonErr.Error.Code = http.StatusBadRequest
	case errors.Is(err, ErrHttpInvalidRequestData):
		jsonErr.Error.Code = http.StatusUnprocessableEntity
	default:
		jsonErr.Error.Code = http.StatusInternalServerError
	}

	return &jsonErr
}
