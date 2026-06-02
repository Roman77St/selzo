package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/Roman77St/selzo/internal/service/auth"
)

type AuthHandler struct {
	authService auth.Service
	logger      *slog.Logger
}

type RegisterRequest struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     user.Role `json:"role"`
}

type RegisterResponse struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Role    user.Role `json:"role"`
	Message string    `json:"message"`
}

func NewAuthHandler(logger *slog.Logger, authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
		logger:      logger,
	}
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}
	err := h.authService.Register(
		r.Context(),
		auth.RegisterUserInput{
			Email:    req.Email,
			Password: req.Password,
			Role:     req.Role,
		},
	)
	if err != nil {
		h.logger.Error("failed to register user", slog.String("error", err.Error()))
		http.Error(
			w,
			"failed to register user",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}
