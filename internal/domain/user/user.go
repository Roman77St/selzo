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
	ID              uuid.UUID
	Email           string
	EmailVerifiedAt *time.Time
	Role            Role
	IsActive        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func New(email string, role Role) (*User, error) {
	switch role {
	case RoleSupplier, RoleBuyer, RoleAdmin:
		// Valid role
	default:
		return nil, ErrInvalidRole
	}

	now := time.Now()

	return &User{
		ID:        uuid.New(),
		Email:     email,
		Role:      role,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
