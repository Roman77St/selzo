package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/usercredential"
	"github.com/jackc/pgx/v5/pgconn"
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

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			switch pgErr.ConstraintName {

			case ConstraintUserCredentialsUserIDFKey:
				return ErrUserNotFound
			}
		}

		return fmt.Errorf(
			"create user credentials: %w",
			err,
		)
	}

	return nil
}
