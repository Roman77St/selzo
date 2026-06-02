package authmodule

import (
	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/repository/postgres"
	"github.com/Roman77St/selzo/internal/service/auth"
	"github.com/Roman77St/selzo/internal/security/password"
)

func New(database *db.DB) *auth.Service {
	userRepo := postgres.NewUserRepository(database)
	credentialRepo := postgres.NewUserCredentialsRepository(database)

	passwordHasher := password.NewArgon2IDHasher()

	return auth.New(
		database,
		userRepo,
		credentialRepo,
		passwordHasher,
	)
}