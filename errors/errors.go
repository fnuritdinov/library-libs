package errors

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrInvalidID      = errors.New("invalid ID")
	ErrAlreadyExist   = errors.New("already exist")
	ErrInvalidInput   = errors.New("invalid input")
	ErrInvalidRequest = errors.New("invalid request")
)
