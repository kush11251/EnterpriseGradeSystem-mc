package errors

import (
	"errors"
)

// NewError returns a new error
(func NewError(message string) error {
	return errors.New(message)
})
