package gameerr

import (
	"context"
	"fmt"
	"log/slog"
)

// ErrorCode Паттерн именования: ErrCode<Домен><Проблема> или просто <ДОМЕН>_<ПРОБЛЕМА>
// Само значение ERR_<ДОМЕН>_<ПРОБЛЕМА>
type ErrorCode string

const (
	ErrCodeStateUnknown ErrorCode = "ERR_STATE_UNKNOWN"

	ErrCodeAssetsLoading ErrorCode = "ERR_ASSETS_LOADING"

	ErrCodeSettingsMarshal   ErrorCode = "ERR_SETTINGS_MARSHAL"
	ErrCodeSettingsUnmarshal ErrorCode = "ERR_SETTINGS_UNMARSHAL"
	ErrCodeSettingsWriting   ErrorCode = "ERR_SETTINGS_WRITEING"
	ErrCodeSettingsReading   ErrorCode = "ERR_SETTINGS_READING"
	ErrCodeSettingsLoading   ErrorCode = "ERR_SETTINGS_LOADING"
	ErrCodeSettingsClosing   ErrorCode = "ERR_SETTINGS_CLOSING"
)

// GameError - кастомная ошибка для игровой логики.
type GameError struct {
	Code    ErrorCode      // Уникальный код ошибки, например: "ERR_ASSET_NOT_FOUND"
	Message string         // Человекочитаемое описание того, что пошло не так на этом уровне
	Err     error          // Оригинальная ошибка, которую мы оборачиваем (может быть nil)
	Meta    map[string]any // Дополнительные контекстные данные (имя файла, ID игрока и т.д.)
}

type GameErrorer interface {
	Error() string
	Unwrap() error
	LogAndReturn(ctx context.Context, level slog.Level) GameErrorer
	Log(ctx context.Context, level slog.Level)
	slog.LogValuer
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
		slog.String("code", string(e.Code)),
	}

	for k, v := range e.Meta {
		attrs = append(attrs, slog.Any(k, v))
	}

	if e.Err != nil {
		attrs = append(attrs, slog.Any("wrapped_error", e.Err))
	}

	return slog.GroupValue(attrs...)
}

func New(code ErrorCode, msg string, err error, meta map[string]any) GameErrorer {
	return &GameError{
		Code:    code,
		Message: msg,
		Err:     err,
		Meta:    meta,
	}
}

func (e *GameError) LogAndReturn(ctx context.Context, level slog.Level) GameErrorer {
	e.Log(ctx, level)
	return e
}

func (e *GameError) Log(ctx context.Context, level slog.Level) {
	slog.LogAttrs(ctx, level, e.Message,
		slog.Any("error_details", e), // вызов LogValue()
	)
}
