package postgres

import (
	"context"
	"fmt"

	"github.com/Roman77St/salzo/internal/db"
	"github.com/Roman77St/salzo/internal/domain/refreshtoken"
)

type RefreshTokenRepository struct {
	db db.DBTX
}


func NewRefreshTokenRepository(
	db db.DBTX,
) *RefreshTokenRepository {

	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(
	ctx context.Context,
	token *refreshtoken.Token,
) error {

	query := `
	INSERT INTO refresh_tokens (
		id,
		user_id,
		token_hash,
		expires_at,
		revoked_at,
		created_at
	)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
		token.RevokedAt,
		token.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf(
			"create refresh token: %w",
			err,
		)
	}

	return nil
}