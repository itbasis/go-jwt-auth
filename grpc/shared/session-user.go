package shared

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/status"
)

func GetSessionUser(ctx context.Context) (*model.SessionUser, *status.Status) {
	logger := zerolog.Ctx(ctx)

	sessionUser, ok := ctx.Value(model.SessionUser{}).(*model.SessionUser)
	logger.Trace().Msgf(model.LogSessionUser, sessionUser)

	if sessionUser == nil {
		errorSessionWithoutAuth := ErrSessionWithoutAuth
		logger.Error().Err(errorSessionWithoutAuth.Err()).Send()

		return nil, errorSessionWithoutAuth
	}

	if !ok || sessionUser.UID == uuid.Nil {
		errorSessionInvalidUser := ErrSessionInvalidUser
		logger.Error().Err(errorSessionInvalidUser.Err()).Msgf(model.LogSessionUser, sessionUser)

		return nil, errorSessionInvalidUser
	}

	return sessionUser, nil
}
