package postgres

import (
	"context"
	"fmt"

	"github.com/Roman77St/selzo/internal/db"
	"github.com/Roman77St/selzo/internal/domain/user"
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

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *user.User,
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
	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.Email,
		user.EmailVerifiedAt,
		user.Role,
		user.IsActive,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}