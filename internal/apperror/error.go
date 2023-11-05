package apperror

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

var (
	ErrURLExist   = NewError("Url already exist", errors.New("url_exist"))
	ErrNotFound   = NewError("Url not found", errors.New("not_found"))
	ErrInvalidURL = NewError("Input url invalid", errors.New("invalid_url"))
)

type AppError struct {
	Msg string `json:"message"`
	Err error  `json:"-"`
}

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %v", a.Msg, a.Err)
}

func NewError(msg string, err error) *AppError {
	return &AppError{
		Msg: msg,
		Err: err,
	}
}

func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, ErrURLExist):
		return codes.AlreadyExists
	case errors.Is(err, ErrNotFound):
		return codes.NotFound
	case errors.Is(err, ErrInvalidURL):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "redis"):
		return codes.NotFound
	}

	return codes.Internal
}

func ParseHTTPErrStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrURLExist):
		return http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrInvalidURL):
		return http.StatusBadRequest
	case strings.Contains(err.Error(), "redis"):
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
