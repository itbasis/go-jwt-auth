//go:generate mockery --all
package jwttoken

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itbasis/go-jwt-auth/model"
)

type JwtToken interface {
	SetSecretKey(secretKey []byte, signMethod jwt.SigningMethod)

	CreateAccessToken(context.Context, model.SessionUser) (string, *time.Time, error)

	Parse(ctx context.Context, tokenString string) (*model.SessionUser, error)
}
