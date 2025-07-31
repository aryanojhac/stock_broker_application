package constants

import "errors"

var (
	ErrUsernameTaken   = errors.New("username already exists")
	ErrEmailTaken      = errors.New("email already exists")
	ErrInsertUser      = errors.New("failed to create user")
	ErrCheckUsername   = errors.New("failed to check username")
	ErrCheckEmail      = errors.New("failed to check email")
	ErrInvalidRequest  = errors.New("invalid request body")
	ErrRegistration    = errors.New("registration failed")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)
