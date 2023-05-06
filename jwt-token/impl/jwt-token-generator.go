package impl

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
)

var ErrCreatingToken = errors.New("error creating token")

func (receiver *JwtTokenImpl) CreateAccessToken(ctx context.Context, sessionUser model.SessionUser) (string, *time.Time, error) {
	return receiver.CreateTokenCustomDuration(ctx, sessionUser, receiver.accessTokenDuration)
}

func (receiver *JwtTokenImpl) CreateRefreshToken(ctx context.Context, sessionUser model.SessionUser) (string, *time.Time, error) {
	return receiver.CreateTokenCustomDuration(ctx, sessionUser, receiver.accessTokenDuration)
}

func (receiver *JwtTokenImpl) CreateTokenCustomDuration(ctx context.Context, sessionUser model.SessionUser, expiredAtDuration time.Duration) (
	string,
	*time.Time,
	error,
) {
	logger := zerolog.Ctx(ctx)

	now := receiver.clock.Now()
	expiredAt := now.Add(expiredAtDuration)
	logger.Debug().Msgf("expiredAt: %s", expiredAt)

	claims := &itbasisJwtToken.SessionUserClaims{
		UID: sessionUser.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    sessionUser.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	logger.Debug().Msgf("claims: %++v", claims)

	token := jwt.NewWithClaims(receiver.signMethod, claims)

	signedString, err := token.SignedString(receiver.signSecretKey)
	if err != nil {
		logger.Error().Err(err).Msg(ErrCreatingToken.Error())

		return "", nil, ErrCreatingToken
	}

	logger.Trace().Msgf("access token: %s", signedString)

	return signedString, &expiredAt, nil
}
