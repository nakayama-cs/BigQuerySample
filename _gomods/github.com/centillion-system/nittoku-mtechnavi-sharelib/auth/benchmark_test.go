package auth_test

import (
	"context"
	"mtechnavi/sharelib/auth"
	"mtechnavi/sharelib/cache"
	"testing"
)

func BenchmarkJWT(b *testing.B) {
	if auth.JWTSignerEmail == "" {
		b.SkipNow()
	}
	ctx := context.Background()

	lru := &cache.LRU{}
	if err := lru.Init(ctx); err != nil {
		panic(err)
	}
	ctx = cache.ContextWithCache(ctx, lru)

	b.Run("SignedJWT", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := auth.SignedJWT(ctx, "a", "b")
			if err != nil {
				b.Error(err)
			}
		}
	})
	b.Run("VerifySignedJWT", func(b *testing.B) {
		b.StopTimer()
		tok, err := auth.SignedJWT(ctx, "a", "b")
		if err != nil {
			b.Log(err)
			b.FailNow()
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, err := auth.VerifySignedJWT(ctx, tok)
			if err != nil {
				b.Error(err)
			}
		}
	})
}
