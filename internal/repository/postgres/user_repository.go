package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db db.DBTX
}

// NewUserRepository creates a new PostgreSQL user repository.
func NewUserRepository(db db.DBTX) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	u *user.User,
) error {
	query := `
		INSERT INTO users (
		id,
		email,
		email_verified_at,
		role, is_active,
		created_at,
		updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	executor := db.DBTXFromContext(ctx, r.db)
	_, err := executor.Exec(
		ctx,
		query,
		u.ID,
		u.Email,
		u.EmailVerifiedAt,
		u.Role,
		u.IsActive,
		u.CreatedAt,
		u.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				switch pgErr.ConstraintName {
				case ConstraintUsersEmailUnique:
					return ErrDuplicateEmail
				}
			}
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}


func (r *UserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	return nil, errors.New("GetUserByEmail not imlemented")
}
