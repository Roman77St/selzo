package auth

import (
	"context"

	"github.com/Roman77St/selzo/internal/domain/user"
)

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user *user.User,
	) error

	CreateCredentials(
		ctx context.Context,
		credentials *user.Credentials,
	) error

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*user.User, error)
}