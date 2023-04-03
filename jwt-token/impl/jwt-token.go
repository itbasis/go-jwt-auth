package impl

import (
	"fmt"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang-jwt/jwt/v5"
	itbasisCoreUtils "github.com/itbasis/go-core-utils"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	"github.com/rs/zerolog/log"
)

//goland:noinspection GoNameStartsWithPackageName
type JwtTokenImpl struct {
	itbasisJwtToken.JwtToken

	clock clock.Clock

	signSecretKey []byte
	signMethod    jwt.SigningMethod

	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJwtToken(clock clock.Clock) (itbasisJwtToken.JwtToken, error) {
	config := itbasisJwtToken.Config{}

	if err := itbasisCoreUtils.ReadEnvConfig(&config); err != nil {
		return nil, err //nolint:wrapcheck
	}

	log.Trace().Msgf("config: %++v", config)

	return NewJwtTokenCustomConfig(clock, config)
}

func NewJwtTokenCustomConfig(clock clock.Clock, config itbasisJwtToken.Config) (itbasisJwtToken.JwtToken, error) {
	obj := &JwtTokenImpl{
		clock:                clock,
		accessTokenDuration:  time.Duration(config.JwtAccessTokenDurationInSeconds),
		refreshTokenDuration: time.Duration(config.JwtRefreshTokenDurationInSeconds),
	}

	if len(config.JwtSecretKey) > 0 {
		signingMethod := jwt.GetSigningMethod(config.JwtSigningMethod)
		log.Info().Msgf("Using signing method: %++v", signingMethod)

		if signingMethod == jwt.SigningMethodNone {
			return nil, fmt.Errorf("%w: %s", jwt.ErrInvalidKeyType, config.JwtSigningMethod)
		}

		obj.SetSecretKey([]byte(config.JwtSecretKey), signingMethod)
	}

	return obj, nil
}

func (receiver *JwtTokenImpl) SetSecretKey(secretKey []byte, signMethod jwt.SigningMethod) {
	receiver.signSecretKey = secretKey
	receiver.signMethod = signMethod
}
