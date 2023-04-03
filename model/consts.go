package model

const HeaderAuthorize = "authorization"
const AuthSchemaBearer string = "bearer"

const LogMdcSessionUserID = "sessionUserUID"

const LogJwtToken = "jwtToken: %++v" //nolint:gosec
const LogAuthUser = "authUser: %++v"
const LogSessionUser = "sessionUser: %++v"

const (
	ErrSessionWithoutAuth     = "session without authentication"
	ErrAuthTokenNotFound      = "authentication token not found"
	ErrAuthenticationRequired = "authentication required"
	ErrSessionInvalidUser     = "session contains an invalid user object"
)
