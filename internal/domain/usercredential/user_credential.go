package usercredential

import (
	"time"

	"github.com/google/uuid"
)

// UserCredential stores authentication data for a user.
type UserCredential struct {
	UserID uuid.UUID

	PasswordHash string

	PasswordChangedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(
	userID uuid.UUID,
	passwordHash string,
) *UserCredential {

	now := time.Now()

	return &UserCredential{
		UserID:            userID,
		PasswordHash:      passwordHash,
		PasswordChangedAt: now,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}
