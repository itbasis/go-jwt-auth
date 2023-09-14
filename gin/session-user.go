package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/itbasis/go-jwt-auth/v2/model"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
)

func GetSessionUser(ctx *gin.Context) (*model.SessionUser, error) {
	logger := zapctx.Logger(ctx).Sugar()

	value, exists := ctx.Get(ctxSessionUser)
	if !exists {
		err := errors.Wrap(model.ErrSessionWithoutAuth, "SessionUser not found in context")
		logger.Error(err)

		return nil, err
	}

	logger.Debugf("found session user: %++v", value)

	sessionUser, ok := value.(*model.SessionUser)
	if !ok {
		err := errors.Wrap(model.ErrSessionInvalidUser, "model cannot be cast to SessionUser")
		logger.Error(err)

		return nil, err
	}

	return sessionUser, nil
}
