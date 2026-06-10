package authmodule

import (
	"time"

	"github.com/Roman77St/salzo/internal/db"
	"github.com/Roman77St/salzo/internal/repository/postgres"
	"github.com/Roman77St/salzo/internal/security/jwt"
	"github.com/Roman77St/salzo/internal/security/password"
	"github.com/Roman77St/salzo/internal/service/auth"
)

func New(database *db.DB, jwtSecret string) *auth.Service {
	userStore := postgres.NewUserRepository(database)
	credentialStore := postgres.NewUserCredentialsRepository(database)
	refreshTokenStore := postgres.NewRefreshTokenRepository(database)

	passwordHasher := password.NewArgon2IDHasher()

	jwtService := jwt.New(
		jwtSecret,
		15*time.Minute,
		30*24*time.Hour,
	)

	return auth.New(
		database,
		userStore,
		credentialStore,
		refreshTokenStore,
		passwordHasher,
		jwtService,
	)
}
