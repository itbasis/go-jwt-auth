package jwttoken

import (
	"strconv"
	"time"
)

const (
	DefaultAccessTokenDuration  = time.Hour * 24
	DefaultRefreshTokenDuration = time.Hour * 24 * 30
)

type Config struct {
	JwtSecretKey     string `env:"JWT_SECRET_KEY"`
	JwtSigningMethod string `env:"JWT_SIGNING_METHOD" envDefault:"HS512"`

	// default: 15 minutes
	JwtAccessTokenDurationInSeconds  TokenDuration `env:"JWT_ACCESS_TOKEN_DURATION" envDefault:"86400"`
	JwtRefreshTokenDurationInSeconds TokenDuration `env:"JWT_REFRESH_TOKEN_DURATION" envDefault:"2592000"`
}

type TokenDuration time.Duration

func (receiver *TokenDuration) UnmarshalText(text []byte) error {
	duration, err := strconv.ParseUint(string(text), 10, 64)

	if err != nil {
		return err //nolint:wrapcheck
	}

	*receiver = TokenDuration(time.Second * time.Duration(duration))

	return nil
}
