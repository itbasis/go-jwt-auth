package server

import (
	"context"

	"github.com/benbjohnson/clock"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	itbasisJwtAuthGrpcShared "github.com/itbasis/go-jwt-auth/grpc/shared"
	itbasisJwtToken "github.com/itbasis/go-jwt-auth/jwt-token"
	itbasisJwtTokenImpl "github.com/itbasis/go-jwt-auth/jwt-token/impl"
	itbasisJwtAuthModel "github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AuthServerInterceptor struct {
	jwtToken itbasisJwtToken.JwtToken
}

func NewAuthServerInterceptor() *AuthServerInterceptor {
	jwtToken, err := itbasisJwtTokenImpl.NewJwtToken(clock.New())
	if err != nil {
		log.Error().Err(err).Msg("")

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
		logger := zerolog.Ctx(ctx)

		headerValue := metadata.ExtractIncoming(ctx).Get(itbasisJwtAuthModel.HeaderAuthorize)
		if headerValue == "" {
			return ctx, nil
		}

		logger.Trace().Msgf("headerValue: %++v", headerValue)

		jwtToken, err := grpcAuth.AuthFromMD(ctx, itbasisJwtAuthModel.AuthSchemaBearer)
		if err != nil {
			logger.Trace().Err(err).Msg("")

			return ctx, err
		}

		logger.Trace().Msgf(itbasisJwtAuthModel.LogJwtToken, jwtToken)

		authUser, err := receiver.jwtToken.Parse(ctx, jwtToken)
		if err != nil {
			logger.Error().Err(err).Msg("")
			// FIXME handle parsing error

			return ctx, itbasisJwtAuthGrpcShared.ErrAuthenticationRequired.Err()
		}

		logger.Trace().Msgf(itbasisJwtAuthModel.LogAuthUser, authUser)

		newCtx := logger.WithContext(context.WithValue(ctx, itbasisJwtAuthModel.SessionUser{}, authUser))

		logger.UpdateContext(
			func(c zerolog.Context) zerolog.Context {
				return c.Str(itbasisJwtAuthModel.LogMdcSessionUserID, authUser.UID.String())
			},
		)

		logger.Debug().Msg("init new context")

		return newCtx, nil
	}
}
