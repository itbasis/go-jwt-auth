package gin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/itbasis/go-clock"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/v2/jwt-token/impl"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"go.uber.org/zap/zapcore"
)

type AuthInterceptor struct {
	jwtToken itbasisJwtToken.JwtToken
}

func NewAuthInterceptor() *AuthInterceptor {
	jwtToken, err := itbasisJwtTokenImpl.NewJwtToken(clock.New())
	if err != nil {
		zapctx.Default.Sugar().Error(err)

		return nil
	}

	return NewAuthInterceptorWithCustomParser(jwtToken)
}

func NewAuthInterceptorWithCustomParser(jwtToken itbasisJwtToken.JwtToken) *AuthInterceptor {
	return &AuthInterceptor{jwtToken: jwtToken}
}

func (receiver *AuthInterceptor) AuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := zapctx.Logger(ctx).Sugar()

		jwtToken, err := receiver.ginGetHeaderAuthorization(ctx)
		if errors.Is(err, model.ErrAuthTokenNotFound) {
			if logger.Level() == zapcore.DebugLevel {
				logger.Error(err)
			}

			ctx.Set(ctxSessionUser, nil)

			return
		}

		logger.Debugf(model.LogJwtToken, jwtToken)

		authUser, err := receiver.jwtToken.Parse(ctx, jwtToken)
		if err != nil {
			logger.Error(err)

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		logger.Debugf(model.LogAuthUser, authUser)

		ctx.Set(ctxSessionUser, authUser)
	}
}

func (receiver *AuthInterceptor) ginGetHeaderAuthorization(ctx *gin.Context) (string, error) {
	logger := zapctx.Logger(ctx).Sugar()

	authHeaderValue := strings.TrimSpace(ctx.GetHeader(model.HeaderAuthorize))
	logger.Debugf("auth header value: %v", authHeaderValue)

	if authHeaderValue == "" {
		err := errors.Wrap(model.ErrAuthTokenNotFound, "header is empty")
		logger.Error(err)

		return "", err
	}

	parts := strings.SplitN(authHeaderValue, " ", tokenParts)

	if len(parts) < tokenParts {
		err := errors.Wrap(itbasisJwtToken.ErrTokenInvalid, fmt.Sprintf("token parts not equal %d", tokenParts))
		logger.Error(err)

		return "", err
	}

	if !strings.EqualFold(parts[0], model.AuthSchemaBearer) {
		err := errors.Wrap(itbasisJwtToken.ErrTokenInvalid, fmt.Sprintf("Request unauthenticated with %s", model.AuthSchemaBearer))
		logger.Error(err)

		return "", err
	}

	token := parts[1]
	logger.Debugf("token: %v", token)

	return token, nil
}
