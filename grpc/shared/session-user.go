package shared

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
)

func GetSessionUser(ctx context.Context) (*model.SessionUser, error) {
	logger := zapctx.Logger(ctx).Sugar()

	sessionUser, ok := ctx.Value(model.SessionUser{}).(*model.SessionUser)
	logger.Debugf(model.LogSessionUser, sessionUser)

	if sessionUser == nil {
		err := errors.Wrap(model.ErrSessionWithoutAuth, "SessionUser is nil")
		logger.Error(err)

		return nil, err
	}

	if !ok {
		err := errors.Wrap(model.ErrSessionInvalidUser, "model cannot be cast to SessionUser")
		logger.Error(err)

		return nil, err
	}

	if !ok || sessionUser.UID == uuid.Nil {
		err := errors.Wrap(model.ErrSessionInvalidUser, "sessionUser UID is nil")
		logger.Error(err)

		return nil, err
	}

	return sessionUser, nil
}
