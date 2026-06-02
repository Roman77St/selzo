package logger

import (
	"log/slog"
	"os"
)

// New creates application logger based on environment.
func New(env string) *slog.Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		// AddSource: env == "dev",
	}

	switch env {
	case "dev":
		// Human-readable logs for local development.
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, opts)
	case "prod":
		opts.Level = slog.LevelInfo
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	return slog.New(handler)
}
