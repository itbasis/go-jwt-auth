package impl

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
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
	logger := zapctx.Logger(ctx).Sugar()

	now := receiver.clock.Now()
	expiredAt := now.Add(expiredAtDuration)
	logger.Debugf("expiredAt: %s", expiredAt)

	claims := &itbasisJwtToken.SessionUserClaims{
		UID: sessionUser.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    sessionUser.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	if len(sessionUser.Email) > 0 {
		claims.Email = sessionUser.Email
	}

	logger.Debugf("claims: %++v", claims)

	token := jwt.NewWithClaims(receiver.signMethod, claims)

	signedString, err := token.SignedString(receiver.signSecretKey)
	if err != nil {
		err = errors.Wrap(ErrCreatingToken, err.Error())
		logger.Error(err)

		return "", nil, err
	}

	logger.Debugf("access token: %s", signedString)

	return signedString, &expiredAt, nil
}
