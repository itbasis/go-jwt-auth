package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/itbasis/go-jwt-auth/model"
	"github.com/rs/zerolog/log"
)

func GetSessionUser(ctx *gin.Context) (*model.SessionUser, error) {
	logger := log.Logger

	value, exists := ctx.Get(ctxSessionUser)
	if !exists {
		logger.Error().Err(ErrorSessionWithoutAuth).Msg("")

		return nil, ErrorSessionWithoutAuth
	}

	logger.Trace().Msgf("found session user: %++v", value)

	sessionUser, ok := value.(*model.SessionUser)
	if !ok {
		logger.Error().Err(ErrorSessionInvalidUser).Msg("")

		return nil, ErrorSessionInvalidUser
	}

	return sessionUser, nil
}
