package server

import (
	"context"

	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/itbasis/go-clock/v2"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/v2/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/v2/jwt-token/impl"
	itbasisJwtAuthModel "github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

type AuthServerInterceptor struct {
	jwtToken itbasisJwtToken.JwtToken
}

func NewAuthServerInterceptor() *AuthServerInterceptor {
	jwtToken, err := itbasisJwtTokenImpl.NewJwtToken(clock.New())
	if err != nil {
		zapctx.Default.Sugar().Error(err)

		return nil
	}

	return NewAuthServerInterceptorWithCustomParser(jwtToken)
}

func NewAuthServerInterceptorWithCustomParser(jwtToken itbasisJwtToken.JwtToken) *AuthServerInterceptor {
	return &AuthServerInterceptor{
		jwtToken: jwtToken,
	}
}

func (receiver *AuthServerInterceptor) GetAuthFunc() grpcAuth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		logger := zapctx.Logger(ctx).Sugar()

		headerValue := metadata.ExtractIncoming(ctx).Get(itbasisJwtAuthModel.HeaderAuthorize)
		if headerValue == "" {
			return ctx, nil
		}

		logger.Debugf("headerValue: %++v", headerValue)

		jwtToken, err := grpcAuth.AuthFromMD(ctx, itbasisJwtAuthModel.AuthSchemaBearer)
		if err != nil {
			return ctx, errors.Wrapf(err, "error getting token")
		}

		logger.Debugf(itbasisJwtAuthModel.LogJwtToken, jwtToken)

		authUser, err := receiver.jwtToken.Parse(ctx, jwtToken)
		if err != nil {
			return ctx, errors.Wrapf(err, "token parsing error")
		}

		logger.Debugf(itbasisJwtAuthModel.LogAuthUser, authUser)

		newCtx := zapctx.WithFields(
			context.WithValue(ctx, itbasisJwtAuthModel.SessionUser{}, authUser),
			zapcore.Field{
				Key:    itbasisJwtAuthModel.LogMdcSessionUserUID,
				Type:   zapcore.StringerType,
				String: authUser.UID.String(),
			},
		)

		return newCtx, nil
	}
}
