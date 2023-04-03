package shared

import (
	"github.com/itbasis/go-jwt-auth/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrSessionWithoutAuth = status.New(codes.Unauthenticated, model.ErrSessionWithoutAuth)
var ErrAuthenticationRequired = status.New(codes.Unauthenticated, model.ErrAuthenticationRequired)
var ErrSessionInvalidUser = status.New(codes.Unauthenticated, model.ErrSessionInvalidUser)
