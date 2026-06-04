package authctx

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/Roman77St/salzo/internal/domain/user"
)

var ErrUserMissingInContext = errors.New("user missing in context")

type User struct {
	ID   uuid.UUID
	Role user.Role
}

type contextKey struct{}

var userContextKey = contextKey{}

func WithUser(
	ctx context.Context,
	user User,
) context.Context {
	return context.WithValue(
		ctx,
		userContextKey,
		user,
	)
}

func UserFromContext(
	ctx context.Context,
) (User, error) {

	user, ok := ctx.Value(userContextKey).(User)
	if !ok {
		return User{}, ErrUserMissingInContext
	}

	return user, nil
}
