package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Roman77St/selzo/internal/handler"
	"github.com/Roman77St/selzo/internal/middleware"
	"github.com/Roman77St/selzo/internal/service/auth"
	"github.com/go-chi/chi/v5"
)

const (
	AuthPrefix    = "/api/v1/auth"
	// UserPrefix    = "/api/v1/users"
	// ProductPrefix = "/api/v1/products"
)

func NewServer(
	addr string,
	logger *slog.Logger,
	authService *auth.Service,
	) *http.Server {

	r := chi.NewRouter()

	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))

	r.Get("/health", handler.HealthHandler)

	r.Route(AuthPrefix, func(auth chi.Router) {
		RegisterAuthRoutes(auth, authService)
	})

	return &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       5 * time.Second,      // 5 seconds
		WriteTimeout:      10 * time.Second,     // 10 seconds
		IdleTimeout:       2 * 60 * time.Second, // 2 minutes
		ReadHeaderTimeout: 2 * time.Second,      // 2 seconds
	}
}
