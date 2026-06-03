package authmodule

import (
	"time"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/repository/postgres"
	"github.com/Roman77St/selzo/internal/security/jwt"
	"github.com/Roman77St/selzo/internal/security/password"
	"github.com/Roman77St/selzo/internal/service/auth"
)

func New(database *db.DB, jwtSecret string) *auth.Service {
	userRepo := postgres.NewUserRepository(database)
	credentialRepo := postgres.NewUserCredentialsRepository(database)

	passwordHasher := password.NewArgon2IDHasher()

	jwtService := jwt.New(
		jwtSecret,
		15*time.Minute,
	)

	return auth.New(
		database,
		userRepo,
		credentialRepo,
		passwordHasher,
		jwtService,
	)
}
