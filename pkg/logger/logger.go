package logger

import (
	"log/slog"
	"os"
)

type contextKey string

const (
	UserIdKey contextKey = "user_id"
	LocaleKey contextKey = "locale"
)

func InitClient(level slog.Level, version string, appName string) {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	baseHandler := slog.NewJSONHandler(os.Stdout, opts)

	clientHandler := &ClientLogHandler{Handler: baseHandler}

	clientLogger := slog.New(clientHandler).With("app", appName, "version", version)

	slog.SetDefault(clientLogger)
}

func InitServer(level slog.Level) {
	opts := &slog.HandlerOptions{Level: level}
	baseHandler := slog.NewJSONHandler(os.Stdout, opts)

	serverHandler := &ServerLogHandler{Handler: baseHandler}

	slog.SetDefault(slog.New(serverHandler))
}
