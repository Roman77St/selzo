package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/usercredential"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

)

// UserCredentialsRepository provides access
// to user authentication credentials storage.
type UserCredentialsRepository struct {
	db db.DBTX
}

// NewUserCredentialsRepository creates a new
// PostgreSQL user credentials repository.
func NewUserCredentialsRepository(
	db db.DBTX,
) *UserCredentialsRepository {

	return &UserCredentialsRepository{
		db: db,
	}
}

// Create inserts new user credentials into the database.
func (r *UserCredentialsRepository) Create(
	ctx context.Context,
	credentials *usercredential.UserCredential,
) error {

	query := `
		INSERT INTO user_credentials (
			user_id,
			password_hash,
			password_changed_at,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`
	executor := db.DBTXFromContext(ctx, r.db)

	_, err := executor.Exec(
		ctx,
		query,
		credentials.UserID,
		credentials.PasswordHash,
		credentials.PasswordChangedAt,
		credentials.CreatedAt,
		credentials.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf(
			"create user credentials: %w",
			mapPostgresError(err),
		)
	}

	return nil
}

func (r *UserCredentialsRepository) GetByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*usercredential.UserCredential, error) {

	query := `
		SELECT
			user_id,
			password_hash,
			password_changed_at,
			created_at,
			updated_at
		FROM user_credentials
		WHERE user_id = $1
	`
	executor := db.DBTXFromContext(ctx, r.db)

	var credentials usercredential.UserCredential

	err := executor.QueryRow(
		ctx,
		query,
		userID,
	).Scan(
		&credentials.UserID,
		&credentials.PasswordHash,
		&credentials.PasswordChangedAt,
		&credentials.CreatedAt,
		&credentials.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserCredentialNotFound
		}

		return nil, fmt.Errorf(
			"get user credentials by user ID: %w",
			err,
		)
	}

	return &credentials, nil
}