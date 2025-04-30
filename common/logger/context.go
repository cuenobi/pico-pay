package logger

import (
	"context"
)

type ctxKey struct{}

func WithRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, ctxKey{}, reqID)
}

func FromContext(ctx context.Context) string {
	reqID, _ := ctx.Value(ctxKey{}).(string)
	return reqID
}

func InfoCtx(ctx context.Context, msg string, fields ...any) {
	reqID := FromContext(ctx)
	Info(msg, append(fields, "request_id", reqID)...)
}
