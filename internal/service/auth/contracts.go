package auth

import (
	"context"

	"github.com/Roman77St/salzo/internal/db"
	"github.com/Roman77St/salzo/internal/domain/refreshtoken"
	"github.com/Roman77St/salzo/internal/domain/user"
	"github.com/Roman77St/salzo/internal/domain/usercredential"
	"github.com/Roman77St/salzo/internal/security/jwt"
	"github.com/google/uuid"
)

// Service provides authentication use cases.
type Service struct {
	db                *db.DB
	userStore         UserStore
	credentialStore   UserCredentialStore
	refreshTokenStore RefreshTokenStore
	passwordHasher    PasswordHasher
	jwtService        *jwt.Service
}

// RegisterUserInput contains data required
// to register a new user.
type RegisterUserInput struct {
	Email    string
	Password string
	Role     user.Role
}

type LoginUserInput struct {
	Email    string
	Password string
}

// UserRepository provides access to user storage.
type UserStore interface {
	Create(
		ctx context.Context,
		user *user.User,
	) error

	GetByEmail(
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

	GetByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*usercredential.UserCredential, error)
}

// PasswordHasher hashes passwords for storage.
type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password, hash string) (bool, error)
}

func (s *Service) ParseToken(
	token string,
) (*jwt.Claims, error) {
	return s.jwtService.Parse(token)
}

type RefreshTokenStore interface {
	Create(
		ctx context.Context,
		token *refreshtoken.Token,
	) error

	// GetByHash(
	// 	ctx context.Context,
	// 	hash string,
	// ) (*refreshtoken.Token, error)

	// Revoke(
	// 	ctx context.Context,
	// 	hash string,
	// ) error
}