package gameerr

import (
	"context"
	"fmt"
	"log/slog"
	"main/pkg/logger"
)

// ErrorCode Паттерн именования: ErrCode<Домен><Проблема> или просто <ДОМЕН>_<ПРОБЛЕМА>
// Само значение ERR_<ДОМЕН>_<ПРОБЛЕМА>
type ErrorCode string

const (
	ErrCodeStateUnknown ErrorCode = "ERR_STATE_UNKNOWN"

	ErrCodeAssetsLoading ErrorCode = "ERR_ASSETS_LOADING"

	ErrCodeSettingsMarshal   ErrorCode = "ERR_SETTINGS_MARSHAL"
	ErrCodeSettingsUnmarshal ErrorCode = "ERR_SETTINGS_UNMARSHAL"
	ErrCodeSettingsWriting   ErrorCode = "ERR_SETTINGS_WRITING"
	ErrCodeSettingsReading   ErrorCode = "ERR_SETTINGS_READING"
	ErrCodeSettingsLoading   ErrorCode = "ERR_SETTINGS_LOADING"
	ErrCodeSettingsClosing   ErrorCode = "ERR_SETTINGS_CLOSING"
)

type ErrorSeverity int8

const (
	SeverityInfo    ErrorSeverity = 1
	SeverityWarning ErrorSeverity = 2
	SeverityFatal   ErrorSeverity = 3
)

// GameError - кастомная ошибка для игровой логики.
type GameError struct {
	Code     ErrorCode
	Message  string
	Err      error
	Meta     map[string]any
	Severity ErrorSeverity
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

func New(code ErrorCode, msg string, err error, serv ErrorSeverity, meta map[string]any) *GameError {
	return &GameError{
		Code:     code,
		Message:  msg,
		Err:      err,
		Severity: serv,
		Meta:     meta,
	}
}

func (e *GameError) Log(ctx context.Context) {
	switch e.Severity {
	case SeverityInfo:
		logger.Info(ctx, e.Message, e.LogValue())
	case SeverityWarning:
		logger.Warn(ctx, e.Message, e.LogValue())
	case SeverityFatal:
		logger.Error(ctx, e.Message, e.LogValue())
	}
}
