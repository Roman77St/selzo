package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Roman77St/salzo/internal/db"
	"github.com/Roman77St/salzo/internal/domain/refreshtoken"
	"github.com/Roman77St/salzo/internal/domain/user"
	"github.com/Roman77St/salzo/internal/domain/usercredential"
	"github.com/Roman77St/salzo/internal/repository/postgres"
	"github.com/Roman77St/salzo/internal/security/jwt"
	"github.com/Roman77St/salzo/internal/security/tokenhash"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken string
	RefreshToken string
}

func New(
	db *db.DB,
	userStore UserStore,
	credentialStore UserCredentialStore,
	refreshTokenStore RefreshTokenStore,
	passwordHasher PasswordHasher,
	jwtService *jwt.Service,
) *Service {
	return &Service{
		db:              db,
		userStore:       userStore,
		credentialStore: credentialStore,
		refreshTokenStore: refreshTokenStore,
		passwordHasher:  passwordHasher,
		jwtService:      jwtService,
	}
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

func (s *Service) Login(
	ctx context.Context,
	input LoginUserInput,
) (*TokenPair, error) {
	user, err := s.userStore.GetByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user: %w", err)
	}

	cred, err := s.credentialStore.GetByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, postgres.ErrUserCredentialNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get credentials: %w", err)
	}

	ok, err := s.passwordHasher.Verify(input.Password, cred.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("verify password: %w", err)
	}
	if !ok {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	tokenHash := tokenhash.Hash(refreshToken)
	rt := refreshtoken.Token{
		ID: uuid.New(),
		UserID: user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(30*24*time.Hour),
		CreatedAt: time.Now(),
	}

	if err := s.refreshTokenStore.Create(ctx, &rt); err != nil {
		return nil, fmt.Errorf(
			"save refresh token: %w",
			err,
		)
	}

	tokenPair := TokenPair{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return &tokenPair, nil
}
