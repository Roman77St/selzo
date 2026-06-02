package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/Roman77St/selzo/internal/handler"
	"github.com/Roman77St/selzo/internal/service/auth"
)

func RegisterAuthRoutes(
	r chi.Router,
	authService *auth.Service,
) {
	authHandler := handler.NewAuthHandler(authService)
	r.Post("/register", authHandler.Register)
}