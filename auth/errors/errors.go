package errors

import "errors"

var (
	//ErrInvalidParameters invalid parameter received
	ErrInvalidParameters = errors.New("Invalid parameters")
	//ErrUnknown unknown error
	ErrUnknown = errors.New("Unknown error")
	//ErrUnauthorized unknown error
	ErrUnauthorized = errors.New("You do not have permission to access this content")
	//ErrUndefinedToken this token was not defined
	ErrUndefinedToken = errors.New("Undefined token")
	//ErrTokenExpired token expired
	ErrTokenExpired = errors.New("Token expired")
	//ErrValidateToken error to validate the token
	ErrValidateToken = errors.New("Error to validate")
	//ErrIDRepeated error to validate the token
	ErrIDRepeated = errors.New("This id already exists")
)
