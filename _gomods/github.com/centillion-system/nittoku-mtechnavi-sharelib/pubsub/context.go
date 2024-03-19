package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc/metadata"
)

func NewAttributeFromContext(ctx context.Context) (map[string]string, error) {
	logger := ctxlogrus.Extract(ctx)

	attr := map[string]string{}

	md, _ := metadata.FromOutgoingContext(ctx)

	var req http.Request
	req.Header = http.Header{
		"Cookie": md.Get("cookie"),
	}
	var cookies []string
	for _, cookie := range req.Cookies() {
		if cookie.Name != "__session" {
			continue
		}
		cookies = append(cookies, cookie.String())
	}
	// 認証に必要なメタデータを付与
	if len(cookies) > 0 {
		if b, err := json.Marshal(cookies); err != nil {
			logger.WithError(err).Error("fail to marshal cookies")
			return nil, err
		} else {
			attr["cookie"] = string(b)
		}
	}
	if xauths := md.Get("x-authorization"); len(xauths) > 0 {
		b, err := json.Marshal(xauths)
		if err != nil {
			logger.WithError(err).Error("fail to marshal x-authorization")
			return nil, err
		}
		// PubSubの制限によりattributeのサイズは1024byteまでのため分割して格納
		// サブスクライバー側で連結して使用する
		// (1024byte以下の場合も分割する場合と同じロジックを通す)
		xAuthz := string(b)
		for i := 0; i*1024 < len(xAuthz); i++ {
			start := i * 1024
			end := 0
			if 1024 > len(xAuthz)-start {
				end = start + (len(xAuthz) - start)
			} else {
				end = start + 1024
			}
			attr[fmt.Sprintf("x-authorization_%02d", i)] = xAuthz[start:end]
		}
	}
	return attr, nil
}
