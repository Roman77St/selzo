package response

import (
	"errors"
	"net/http"

	"github.com/Roman77St/selzo/internal/service/auth"
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
			Status: http.StatusUnauthorized,
			Code: "INVALID_CREDENTIALS",
			Message: "invalid email or password",
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