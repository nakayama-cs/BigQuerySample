package cache

import (
	"context"
	"os"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type ctxKey string

const (
	ctxKeyCache   = ctxKey("cache")
	ctxKeyDisable = ctxKey("disable")
)

var (
	// 全体一括でキャッシュ動作をOffにするフラグ用環境変数名
	// 変数名は外から上書き可能にする
	EnvDisabledCache = "CACHE_DISABLED"
)

func LRUFromContext(ctx context.Context) *LRU {
	v, ok := ctx.Value(ctxKeyCache).(*LRU)
	if !ok {
		v = &LRU{}
		if err := v.Init(ctx); err != nil {
			panic(err)
		}
		return v
	}
	return v
}

func CacheFromContext(ctx context.Context) Cache {
	return LRUFromContext(ctx)
}

func ContextWithCache(parent context.Context, v *LRU) context.Context {
	return context.WithValue(parent, ctxKeyCache, v)
}

func Enabled(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyDisable, false)
}

func Disabled(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyDisable, true)
}

func disabledFromContext(ctx context.Context) bool {
	// 環境変数で指定されたものが最優先
	if s := os.Getenv(EnvDisabledCache); s == "true" {
		return true
	}
	v, ok := ctx.Value(ctxKeyDisable).(bool)
	if !ok {
		return false
	}
	return v
}

func (c *LRU) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ContextWithCache(stream.Context(), c)
		wrapped := grpcmiddleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (c *LRU) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = ContextWithCache(ctx, c)
		return handler(ctx, req)
	}
}
