package http

import (
	"log/slog"

	"github.com/Roman77St/selzo/internal/handler"
	"github.com/Roman77St/selzo/internal/middleware"
	"github.com/Roman77St/selzo/internal/service/auth"
	"github.com/go-chi/chi/v5"
)

// Main route prefixes
const (
	AuthPrefix    = "/api/v1/auth"
)

// newMainRouter sets up the main router with all routes and middleware.
func newRouter(
	authService *auth.Service,
	logger *slog.Logger,
) chi.Router {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))

	// Health check endpoint
	r.Get("/health", handler.HealthHandler)

	r.Route(AuthPrefix, func(auth chi.Router) {
		RegisterAuthRoutes(auth, authService)
	})

	return r
}