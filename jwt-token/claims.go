package jwttoken

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

type SessionUserClaims struct {
	UID   uuid.UUID `json:"uid,omitempty"`
	Email string    `json:"email,omitempty"`

	jwt.RegisteredClaims
}
