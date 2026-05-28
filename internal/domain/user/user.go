package user

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleSupplier Role = "supplier"
	RoleBuyer    Role = "buyer"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID                uuid.UUID
	Email             string
	EmailVerifiedAt   *time.Time
	Role              Role
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}