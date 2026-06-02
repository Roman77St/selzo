package auth

import (
	"context"

	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/Roman77St/selzo/internal/domain/usercredential"
)

// UserRepository provides access to user storage.
type UserStore interface {
	Create(
		ctx context.Context,
		user *user.User,
	) error

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*user.User, error)
}

// UserCredentialRepository provides access
// to user authentication credentials storage.
type UserCredentialStore interface {
	Create(
		ctx context.Context,
		credential *usercredential.UserCredential,
	) error
}

// PasswordHasher hashes passwords for storage.
type PasswordHasher interface {
	Hash(password string) (string, error)
}
