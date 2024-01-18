package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	clock "github.com/itbasis/go-clock/v2/pkg"
	itbasisCoreUtilsEnvReader "github.com/itbasis/go-core-utils/v2/env-reader"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	"github.com/juju/zaputil/zapctx"
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

	if err := itbasisCoreUtilsEnvReader.ReadEnvConfig(context.Background(), &config, nil); err != nil {
		return nil, err //nolint:wrapcheck
	}

	zapctx.Default.Sugar().Debugf("config: %++v", config)

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
		zapctx.Default.Sugar().Infof("Using signing method: %++v", signingMethod)

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
