package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/jackc/pgx/v5"
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
		return fmt.Errorf("create user: %w", mapPostgresError(err))
	}
	return nil
}

func (r *UserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	query := `
		SELECT id, email, email_verified_at, role, is_active, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	executor := db.DBTXFromContext(ctx, r.db)
	row := executor.QueryRow(ctx, query, email)

	var u user.User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.EmailVerifiedAt,
		&u.Role,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &u, nil
}
