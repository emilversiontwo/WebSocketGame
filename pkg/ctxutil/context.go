package ctxutil

import (
	"context"
	"main/pkg/logger"
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, logger.UserIdKey, userID)
}
