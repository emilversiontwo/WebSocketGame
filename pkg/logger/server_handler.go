package logger

import (
	"context"
	"log/slog"
)

type ServerLogHandler struct {
	slog.Handler
}

func (h *ServerLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if v := ctx.Value(UserIdKey); v != nil {
		r.AddAttrs(slog.String(string(UserIdKey), v.(string)))
	}

	return h.Handler.Handle(ctx, r)
}

func (h *ServerLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ServerLogHandler{h.Handler.WithAttrs(attrs)}
}

func (h *ServerLogHandler) WithGroup(name string) slog.Handler {
	return &ServerLogHandler{h.Handler.WithGroup(name)}
}
