package user

import "time"

type Role string

const (
	RoleSupplier Role = "supplier"
	RoleBuyer    Role = "buyer"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}