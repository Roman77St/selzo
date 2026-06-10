package refreshtoken

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID uuid.UUID

	UserID uuid.UUID

	TokenHash string

	ExpiresAt time.Time

	RevokedAt *time.Time

	CreatedAt time.Time
}