package gameerr

import (
	"context"
	"fmt"
	"log/slog"
)

// GameError - кастомная ошибка для игровой логики.
type GameError struct {
	Code    string         // Уникальный код ошибки, например: "ERR_ASSET_NOT_FOUND"
	Message string         // Человекочитаемое описание того, что пошло не так на этом уровне
	Err     error          // Оригинальная ошибка, которую мы оборачиваем (может быть nil)
	Meta    map[string]any // Дополнительные контекстные данные (имя файла, ID игрока и т.д.)
}

func (e *GameError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *GameError) Unwrap() error {
	return e.Err
}

func (e *GameError) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.String("code", e.Code),
	}

	for k, v := range e.Meta {
		attrs = append(attrs, slog.Any(k, v))
	}

	if e.Err != nil {
		attrs = append(attrs, slog.Any("wrapped_error", e.Err))
	}

	return slog.GroupValue(attrs...)
}

func New(code, msg string, err error, meta map[string]any) *GameError {
	return &GameError{
		Code:    code,
		Message: msg,
		Err:     err,
		Meta:    meta,
	}
}

func (e *GameError) LogAndReturn(ctx context.Context, level slog.Level) error {
	e.Log(ctx, level)
	return e
}

func (e *GameError) Log(ctx context.Context, level slog.Level) {
	slog.LogAttrs(ctx, level, e.Message,
		slog.Any("error_details", e), // вызов LogValue()
	)
}
