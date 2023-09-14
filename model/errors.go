package model

import "errors"

var (
	ErrSessionWithoutAuth     = errors.New("session without authentication")
	ErrAuthTokenNotFound      = errors.New("authentication token not found")
	ErrAuthenticationRequired = errors.New("authentication required")
	ErrSessionInvalidUser     = errors.New("session contains an invalid user object")
)
