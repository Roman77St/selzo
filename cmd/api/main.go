package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Roman77St/selzo/internal/config"
	"github.com/Roman77St/selzo/internal/db"
	httpserver "github.com/Roman77St/selzo/internal/http"
	"github.com/Roman77St/selzo/internal/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	logg := logger.New(cfg.AppEnv)

	logg.Info("configuration loaded successfully")
	logg.Info("logger initialized", "environment", cfg.AppEnv)

	pool, err := db.NewPostgresPool(ctx, cfg)
	if err != nil {
		logg.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	logg.Info("postgreSQL database connected")

	server := httpserver.NewServer(
		fmt.Sprintf(":%s", cfg.AppPort),
		logg,
	)

	go func() {
		logg.Info("starting HTTP server", "port", cfg.AppPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("HTTP server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logg.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logg.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	logg.Info("server shutdown complete")
}