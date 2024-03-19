package auth

import (
	"context"
)

type ctxKey string

const (
	ctxKeyInvoker = ctxKey("invoker")
)

func InvokerFromContext(ctx context.Context) *Invoker {
	return ctx.Value(ctxKeyInvoker).(*Invoker)
}

// for testing
func ContextWithInvoker(parent context.Context, v *Invoker) context.Context {
	return context.WithValue(parent, ctxKeyInvoker, v)
}
