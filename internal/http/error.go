package http

import "github.com/dddsphere/quanta/internal/core/errors"

type (
	HTTPError struct {
		errors.Err
	}
)

var (
	CommandDispatchError = NewError("cannot dispatch command")
	QueryDispatchError   = NewError("cannot dispatch query")
)

func NewError(msg string) HTTPError {
	return HTTPError{
		Err: errors.NewError(msg),
	}
}
