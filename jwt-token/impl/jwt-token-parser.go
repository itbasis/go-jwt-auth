package impl

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
)

func (receiver *JwtTokenImpl) Parse(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	logger := zerolog.Ctx(ctx)

	logger.Trace().Msgf("tokenString: %s", tokenString)

	// TODO Adding Firebase parsing

	return receiver.parseWithSecretKey(ctx, tokenString)
}

func (receiver *JwtTokenImpl) parseWithSecretKey(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	logger := zerolog.Ctx(ctx)

	logger.Trace().Msgf("tokenString: %s", tokenString)

	token, err := jwt.ParseWithClaims(
		tokenString, &itbasisJwtToken.SessionUserClaims{}, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				alg := token.Header["alg"]
				logger.Trace().Msgf("alg: %s", alg)

				logger.Error().Err(itbasisJwtToken.ErrUnsupportedSigningMethod).Msgf("token signing algorithm: %v", alg)

				return nil, fmt.Errorf("%w: %v", itbasisJwtToken.ErrUnsupportedSigningMethod, alg)
			}

			logger.Debug().Msgf("method: %v", method)
			logger.Trace().Msgf("using signSecretKey: %s", string(receiver.signSecretKey))

			return receiver.signSecretKey, nil
		},
	)

	if err != nil {
		logger.Error().Err(err).Msgf("")

		return nil, err
	}

	// TODO check - this seems like a redundant check
	if !token.Valid {
		logger.Error().Err(jwt.ErrTokenUnverifiable).Msg("")

		return nil, jwt.ErrTokenUnverifiable
	}

	claims, ok := token.Claims.(*itbasisJwtToken.SessionUserClaims)
	if !ok {
		logger.Error().Err(jwt.ErrTokenInvalidClaims).Msgf("")

		return nil, jwt.ErrTokenInvalidClaims
	}

	logger.Debug().Msgf("claims: %++v", claims)

	sessionUser := &model.SessionUser{}

	if !claims.UID.IsNil() {
		sessionUser.UID = claims.UID
	} else {
		logger.Error().Err(itbasisJwtToken.ErrTokenInvalidUID).Msgf("claims: %++v", claims)
	}

	if len(claims.Issuer) > 0 {
		sessionUser.Username = claims.Issuer
	} else {
		logger.Error().Err(jwt.ErrTokenInvalidIssuer).Msgf("claims: %++v", claims)

		return nil, jwt.ErrTokenInvalidIssuer
	}

	if len(claims.Email) > 0 {
		sessionUser.Email = claims.Email
	} else {
		logger.Error().Err(itbasisJwtToken.ErrTokenInvalidEmail).Msgf("claims: %++v", claims)
	}

	// TODO sessionUser.hasGuest
	// TODO sessionUser.UID

	return sessionUser, nil
}
