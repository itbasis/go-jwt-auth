package impl

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (receiver *JwtTokenImpl) Parse(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	logger := zapctx.Logger(ctx).Sugar()

	logger.Debugf("tokenString: %s", tokenString)

	// TODO Adding Firebase parsing

	return receiver.parseWithSecretKey(ctx, tokenString)
}

func (receiver *JwtTokenImpl) parseWithSecretKey(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	logger := zapctx.Logger(ctx).Sugar()

	logger.Debugf("tokenString: %s", tokenString)

	token, err := jwt.ParseWithClaims(
		tokenString, &itbasisJwtToken.SessionUserClaims{}, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				alg := token.Header["alg"]
				logger.Debugf("alg: %s", alg)

				err := errors.Wrap(itbasisJwtToken.ErrUnsupportedSigningMethod, fmt.Sprintf("token signing algorithm: %v", alg))
				logger.Error(err)

				return nil, err
			}

			logger.Debugf("method: %v", method)
			logger.Debugf("using signSecretKey: %s", string(receiver.signSecretKey))

			return receiver.signSecretKey, nil
		},
	)

	if err != nil {
		err = errors.Wrap(itbasisJwtToken.ErrParsingClaims, err.Error())
		logger.Error(err)

		return nil, err
	}

	// TODO check - this seems like a redundant check
	if !token.Valid {
		logger.Error(jwt.ErrTokenUnverifiable)

		return nil, jwt.ErrTokenUnverifiable
	}

	claims, ok := token.Claims.(*itbasisJwtToken.SessionUserClaims)
	if !ok {
		logger.Error(jwt.ErrTokenInvalidClaims)

		return nil, jwt.ErrTokenInvalidClaims
	}

	logger.Debugf("claims: %++v", claims)

	return receiver.enrichSessionUser(logger, claims)
}

func (receiver *JwtTokenImpl) enrichSessionUser(logger *zap.SugaredLogger, claims *itbasisJwtToken.SessionUserClaims) (
	sessionUser *model.SessionUser,
	err error,
) {
	logger.Debugf("claims: %++v", claims)

	sessionUser = &model.SessionUser{}

	if !claims.UID.IsNil() {
		sessionUser.UID = claims.UID
	} else {
		err = errors.Wrap(itbasisJwtToken.ErrTokenInvalidUID, "UID is nil")
		logger.Error(err)

		return nil, err
	}

	if len(claims.Issuer) > 0 {
		sessionUser.Username = claims.Issuer
	} else {
		err = errors.Wrap(jwt.ErrTokenInvalidIssuer, "is empty")
		logger.Error(err)

		return nil, err
	}

	if len(claims.Email) > 0 {
		sessionUser.Email = claims.Email
	} else {
		err = errors.Wrap(itbasisJwtToken.ErrTokenInvalidEmail, "is empty")
		logger.Error(err)

		return nil, err
	}

	// TODO sessionUser.hasGuest
	// TODO sessionUser.UID

	return sessionUser, nil
}
