package gin

import (
	"errors"
	"net/http"
	"strings"

	"github.com/benbjohnson/clock"
	"github.com/gin-gonic/gin"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/jwt-token/impl"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AuthInterceptor struct {
	jwtToken itbasisJwtToken.JwtToken
}

func NewAuthInterceptor() *AuthInterceptor {
	jwtToken, err := itbasisJwtTokenImpl.NewJwtToken(clock.New())
	if err != nil {
		log.Error().Err(err).Send()

		return nil
	}

	return NewAuthInterceptorWithCustomParser(jwtToken)
}

func NewAuthInterceptorWithCustomParser(jwtToken itbasisJwtToken.JwtToken) *AuthInterceptor {
	return &AuthInterceptor{jwtToken: jwtToken}
}

func (receiver *AuthInterceptor) AuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := zerolog.Ctx(ctx)

		jwtToken, err := receiver.ginGetHeaderAuthorization(ctx)
		if errors.Is(err, ErrorAuthTokenNotFound) {
			logger.Trace().Err(err).Send()

			ctx.Set(ctxSessionUser, nil)

			return
		}

		logger.Trace().Msgf(model.LogJwtToken, jwtToken)

		authUser, err := receiver.jwtToken.Parse(ctx, jwtToken)
		if err != nil {
			logger.Error().Err(err).Send()

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		logger.Trace().Msgf(model.LogAuthUser, authUser)

		ctx.Set(ctxSessionUser, authUser)
	}
}

func (receiver *AuthInterceptor) ginGetHeaderAuthorization(ctx *gin.Context) (string, error) {
	logger := zerolog.Ctx(ctx)

	authHeaderValue := strings.TrimSpace(ctx.GetHeader(model.HeaderAuthorize))
	logger.Trace().Msgf("auth header value: %v", authHeaderValue)

	if authHeaderValue == "" {
		logger.Trace().Err(ErrorAuthTokenNotFound).Send()

		return "", ErrorAuthTokenNotFound
	}

	parts := strings.SplitN(authHeaderValue, " ", 2)

	if len(parts) < 2 {
		logger.Error().Err(itbasisJwtToken.ErrTokenInvalid).Msg("Token parts not equal 2")

		return "", itbasisJwtToken.ErrTokenInvalid
	}

	if !strings.EqualFold(parts[0], model.AuthSchemaBearer) {
		logger.Error().Err(itbasisJwtToken.ErrTokenInvalid).Msgf("Request unauthenticated with %s", model.AuthSchemaBearer)

		return "", itbasisJwtToken.ErrTokenInvalid
	}

	token := parts[1]
	logger.Trace().Msgf("token: %v", token)

	return token, nil
}
