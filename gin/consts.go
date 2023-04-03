package gin

import (
	"errors"

	"github.com/itbasis/go-jwt-auth/model"
)

const ctxSessionUser = "sessionUser"

var ErrorSessionWithoutAuth = errors.New(model.ErrSessionWithoutAuth)
var ErrorAuthTokenNotFound = errors.New(model.ErrAuthTokenNotFound)
var ErrorSessionInvalidUser = errors.New(model.ErrSessionInvalidUser)
