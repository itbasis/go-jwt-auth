package model

const HeaderAuthorize = "authorization"
const AuthSchemaBearer string = "bearer"

const LogMdcSessionUserUID = "sessionUserUID"

// #nosec
const LogJwtToken = "jwtToken: %++v" //nolint:gosec
const LogAuthUser = "authUser: %++v"
const LogSessionUser = "sessionUser: %++v"
