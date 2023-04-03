package jwttoken

import "errors"

var (
	ErrUnsupportedSigningMethod = errors.New("unsupported signing method")
	ErrTokenInvalid             = errors.New("the token has an invalid")

	ErrTokenInvalidUID   = errors.New("token has invalid UID")
	ErrTokenInvalidEmail = errors.New("token has invalid email")
)
