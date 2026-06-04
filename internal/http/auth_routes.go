package http

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/Roman77St/salzo/internal/handler"
	"github.com/Roman77St/salzo/internal/middleware"
	"github.com/Roman77St/salzo/internal/service/auth"
)

func RegisterAuthRoutes(
	r chi.Router,
	authService *auth.Service,
	logger *slog.Logger,
) {
	authHandler := handler.NewAuthHandler(logger, authService)
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.With(middleware.Auth(authService)).Get("/me", authHandler.Me)
}
