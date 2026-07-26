package logger

import (
	"context"
	"log/slog"
	"sync/atomic"
)

var currentClientState atomic.Value

func init() {
	currentClientState.Store("unknown")
}

func SetClientState(state string) {
	currentClientState.Store(state)
}

type ClientLogHandler struct {
	slog.Handler
}

func (h *ClientLogHandler) Handle(ctx context.Context, r slog.Record) error {
	state := currentClientState.Load().(string)
	r.AddAttrs(slog.String("state", state))

	if v := ctx.Value(LocaleKey); v != nil {
		r.AddAttrs(slog.String("locale", v.(string)))
	}

	if v := ctx.Value(UserIdKey); v != nil {
		r.AddAttrs(slog.String(string(UserIdKey), v.(string)))
	}

	return h.Handler.Handle(ctx, r)
}

func (h *ClientLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ClientLogHandler{h.Handler.WithAttrs(attrs)}
}

func (h *ClientLogHandler) WithGroup(name string) slog.Handler {
	return &ClientLogHandler{h.Handler.WithGroup(name)}
}
