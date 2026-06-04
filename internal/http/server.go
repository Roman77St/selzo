package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Roman77St/salzo/internal/service/auth"
)

// NewServer creates and configures the HTTP server,
// registers routes and middleware, and returns
// a ready-to-run http.Server instance.
func NewServer(
	addr string,
	authService *auth.Service,
	logger *slog.Logger,
) *http.Server {

	r := newRouter(authService, logger)

	return &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       5 * time.Second,      // 5 seconds
		WriteTimeout:      10 * time.Second,     // 10 seconds
		IdleTimeout:       2 * 60 * time.Second, // 2 minutes
		ReadHeaderTimeout: 2 * time.Second,      // 2 seconds
	}
}
