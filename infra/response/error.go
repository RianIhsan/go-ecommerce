package response

import "errors"

// General errors
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalid  = errors.New("password must be at least 6 characters")
	ErrAuthIsNotExists  = errors.New("auth is not exists")
	ErrEmailAlreadyUsed = errors.New("email already used")
)
