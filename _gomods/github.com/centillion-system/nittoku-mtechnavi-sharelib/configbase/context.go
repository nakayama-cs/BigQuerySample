package configbase

import (
	"context"
)

type ctxKey string

const (
	ctxKeyPath = ctxKey("path")
)

func pathFromContext(ctx context.Context) string {
	v, _ := ctx.Value(ctxKeyPath).(string)
	return v
}

func contextWithPath(parent context.Context, p string) context.Context {
	return context.WithValue(parent, ctxKeyPath, p)
}
