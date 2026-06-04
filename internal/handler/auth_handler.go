package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Roman77St/selzo/internal/authctx"
	"github.com/Roman77St/selzo/internal/domain/user"
	"github.com/Roman77St/selzo/internal/http/response"
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
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
		h.logger.Error(
			"failed to register user",
			"error", err,
		)
		response.WriteAppError(w, err)

		return
	}

	response.WriteJSON(w, http.StatusCreated, map[string]string{
		"status": "ok",
	}, "USER_REGISTERED")
}

// POST /api/v1/auth/login
func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, "invalid request body", "INVALID_REQUEST")
	}

	token, err := h.authService.Login(
		r.Context(),
		auth.LoginUserInput{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, LoginResponse{
		Token: token,
	}, "OK")
}

func (u *AuthHandler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {
	authUser, err := authctx.UserFromContext(
		r.Context(),
	)

	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	fmt.Println(authUser.ID)
	fmt.Println(authUser.Role)
}