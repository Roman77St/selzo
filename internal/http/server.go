package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Roman77St/selzo/internal/handler"
	"github.com/Roman77St/selzo/internal/middleware"
)

func NewServer(addr string, logger *slog.Logger) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)

	// Wrap the mux with middleware.
	var handler http.Handler = mux

	handler = middleware.LoggingMiddleware(logger, handler)
	handler = middleware.RecoveryMiddleware(logger, handler)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
		ReadTimeout:       5 * time.Second, // 5 seconds
		WriteTimeout:      10 * time.Second, // 10 seconds
		IdleTimeout:       2 * 60 * time.Second, // 2 minutes
		ReadHeaderTimeout: 2 * time.Second, // 2 seconds
	}
}