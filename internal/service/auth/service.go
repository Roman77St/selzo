package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/Roman77St/selzo/internal/domain/usercredential"
	"github.com/Roman77St/selzo/internal/repository/postgres"
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
		db:              db,
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

func (s *Service) Register(
	ctx context.Context,
	input RegisterUserInput,
) error {

	// создаём доменного пользователя
	newUser, err := user.New(input.Email, input.Role)
	if err != nil {
		return fmt.Errorf("create user domain: %w", err)
	}

	// хешируем пароль
	passwordHash, err := s.passwordHasher.Hash(input.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	// создаём credentials
	cred := usercredential.New(newUser.ID, passwordHash)

	err = s.db.WithinTransaction(ctx, func(txCtx context.Context) error {

		// сохраняем пользователя
		if err := s.userStore.Create(txCtx, newUser); err != nil {
			if errors.Is(err, postgres.ErrDuplicateEmail) {
				return ErrUserAlreadyExists
			}
			return fmt.Errorf("create user: %w", err)
		}

		// сохраняем credentials
		if err := s.credentialStore.Create(txCtx, cred); err != nil {
			return fmt.Errorf("create credentials: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("register user transaction: %w", err)
	}

	return nil
}
