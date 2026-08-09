package logger

import (
	"context"
	"log/slog"
	"os"
)

type contextKey string

const (
	UserIdKey contextKey = "user_id"
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

func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}
