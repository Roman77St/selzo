package response

import (
	"errors"
	"net/http"

	"github.com/Roman77St/salzo/internal/authctx"
	"github.com/Roman77St/salzo/internal/security/jwt"
	"github.com/Roman77St/salzo/internal/service/auth"
)

type AppError struct {
	Status  int
	Code    string
	Message string
}

func MapError(err error) AppError {

	switch {

	case errors.Is(err, auth.ErrUserAlreadyExists):
		return AppError{
			Status:  http.StatusConflict,
			Code:    "USER_ALREADY_EXISTS",
			Message: "user with this email already exists",
		}
	case errors.Is(err, auth.ErrInvalidCredentials):
		return AppError{
			Status:  http.StatusUnauthorized,
			Code:    "INVALID_CREDENTIALS",
			Message: "invalid email or password",
		}
	case errors.Is(err, authctx.ErrUserMissingInContext):
		return AppError{
			Status:  http.StatusUnauthorized,
			Code:    "USER_MISSING_IN_CONTEXT",
			Message: "user not found in context",
		}
	case errors.Is(err, auth.ErrUnauthorized):
		return AppError{
			Status:  http.StatusUnauthorized,
			Code:    "UNAUTHORIZED",
			Message: "unauthorized",
		}
	case errors.Is(err, jwt.ErrInvalidToken):
		return AppError{
			Status:  http.StatusUnauthorized,
			Code:    "INVALID_TOKEN",
			Message: "invalid token",
		}
	case errors.Is(err, jwt.ErrTokenExpired):
		return AppError{
			Status:  http.StatusUnauthorized,
			Code:    "TOKEN_EXPIRED",
			Message: "token expired",
		}

	default:
		return AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_ERROR",
			Message: "internal server error",
		}
	}
}

func WriteAppError(
	w http.ResponseWriter,
	err error,
) {
	appErr := MapError(err)

	WriteError(
		w,
		appErr.Status,
		appErr.Message,
		appErr.Code,
	)
}
