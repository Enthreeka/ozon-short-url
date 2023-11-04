package apperror

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
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
	}

	return codes.Internal
}
