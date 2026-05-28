package user

import (
	"time"

	"github.com/google/uuid"
)

type Credentials struct {
	UserID            uuid.UUID
	PasswordHash      string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}