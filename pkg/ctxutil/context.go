package ctxutil

import (
	"context"
	"main/pkg/logger"
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, logger.UserIdKey, userID)
}

// WithLocale метод не имеет смысла тк locale может изменятся и не представляет какой либо ценнсоти для ctx
// TODO: Будет вырезан
func WithLocale(ctx context.Context, locale string) context.Context {
	return context.WithValue(ctx, logger.LocaleKey, locale)
}
