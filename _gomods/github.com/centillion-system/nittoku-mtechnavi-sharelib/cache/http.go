package cache

import (
	"bufio"
	"bytes"
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

// Google Cloudに対し、デフォルトスコープでの認証つきのHTTPClientを返却する。
func (c *LRU) HTTPClient(ctx context.Context, scopes ...string) *http.Client {
	if len(scopes) == 0 {
		scopes = append(scopes, "https://www.googleapis.com/auth/cloud-platform")
	}
	return c.NewHTTPClient(ctx, option.WithScopes(scopes...))
}

func (c *LRU) NewHTTPClient(ctx context.Context, opts ...option.ClientOption) *http.Client {
	hc, _, err := transport.NewHTTPClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	base := hc.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	hc.Transport = &roundTripper{
		Base:  base,
		Cache: c,
	}
	return hc
}

type roundTripper struct {
	Base http.RoundTripper

	Cache *LRU
}

func (tp *roundTripper) getCacheKey(req *http.Request) string {
	return req.Method + " " + req.URL.String()
}

func (tp *roundTripper) getCached(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	if disabledFromContext(ctx) {
		return nil, nil
	}
	cached, ok := tp.Cache.store.Get(tp.getCacheKey(req))
	if !ok {
		return nil, nil
	}
	data, ok := cached.([]byte)
	if !ok {
		return nil, nil
	}
	return http.ReadResponse(bufio.NewReader(bytes.NewReader(data)), req)
}

func (tp *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	logger := ctxlogrus.Extract(ctx)
	cacheKey := tp.getCacheKey(req)
	res, err := tp.getCached(req)
	if res != nil && err == nil {
		// use cached
		logger.Debugf("cache hit: %v", cacheKey)
		return res, nil
	} else if err != nil {
		// go through
		logger.WithError(err).Warnf("failed to get cached http response: %v", cacheKey)
	}

	logger.Debugf("cache missed: %v", cacheKey)
	res, err = tp.Base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	data, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}
	tp.Cache.store.Set(cacheKey, data)
	return res, nil
}
