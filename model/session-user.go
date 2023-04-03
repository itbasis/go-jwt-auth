package model

import "github.com/gofrs/uuid/v5"

type SessionUser struct {
	UID uuid.UUID

	Username string
	Email    string

	HasGuest bool
}
