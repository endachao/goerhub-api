package e

import "errors"

var (
	ErrEmptyAuthHeader         = errors.New("auth header is empty")
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
	ErrExpiredToken            = errors.New("token is expired")
	ErrFailedAuthentication    = errors.New("incorrect Username or Password")
	ErrForbidden               = errors.New("you don't have permission to access this resource")
)
