package http

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/Roman77St/selzo/internal/handler"
	"github.com/Roman77St/selzo/internal/service/auth"
)

func RegisterAuthRoutes(
	r chi.Router,
	authService *auth.Service,
	logger *slog.Logger,
) {
	registerHandler := handler.NewAuthHandler(logger, authService)
	r.Post("/register", registerHandler.Register)
	
	loginHandler := handler.NewAuthHandler(logger, authService)
	r.Post("/login", loginHandler.Login)
}
