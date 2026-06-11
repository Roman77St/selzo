package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roman77St/salzo/internal/db"
	"github.com/Roman77St/salzo/internal/domain/refreshtoken"
	"github.com/jackc/pgx/v5"
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

func (r *RefreshTokenRepository) GetByHash(
	ctx context.Context,
	hash string,
) (*refreshtoken.Token, error) {

	query := `
	SELECT
		id,
		user_id,
		token_hash,
		expires_at,
		revoked_at,
		created_at
	FROM refresh_tokens
	WHERE token_hash = $1
	`

	var token refreshtoken.Token

	err := r.db.QueryRow(
		ctx,
		query,
		hash,
	).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.RevokedAt,
		&token.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRefreshTokenNotFound
		}

		return nil, fmt.Errorf(
			"get refresh token: %w",
			err,
		)
	}

	return &token, nil
}