package exception

import (
	"errors"
)

type Exception struct {
	err error
}

type ExceptionContract interface {
	error
}

// New creates a custom error with the appropriate message.
func New(message string) ExceptionContract {
	return &Exception{
		err: errors.New(message),
	}
}

// Wrap will wraps custom error
func Wrap(cause error) ExceptionContract {
	if _, ok := cause.(*Exception); !ok {
		return &Exception{
			err: cause,
		}
	}

	return cause.(*Exception)
}

// Error will generate error message.
func (e *Exception) Error() string {
	return e.err.Error()
}
