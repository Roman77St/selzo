package auth

import (
	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/user"
)

// Service provides authentication use cases.
type Service struct {
	db              *db.DB
	userStore       UserStore
	credentialStore UserCredentialStore
	passwordHasher  PasswordHasher
}

func New(
	db *db.DB,
	userStore UserStore,
	credentialStore UserCredentialStore,
	passwordHasher PasswordHasher,
) *Service {
	return &Service{
		db:                db,
		userStore:       userStore,
		credentialStore: credentialStore,
		passwordHasher:  passwordHasher,
	}
}

// RegisterUserInput contains data required
// to register a new user.
type RegisterUserInput struct {
	Email    string
	Password string
	Role     user.Role
}
