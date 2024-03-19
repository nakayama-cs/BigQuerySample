package startup

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
	"google.golang.org/api/pubsub/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	EnableTracing bool
)

type ctxKey string

const (
	ctxKeyClientConn = ctxKey("*grpc.ClientConn")
)

func clientConnFromContext(ctx context.Context) *grpc.ClientConn {
	return ctx.Value(ctxKeyClientConn).(*grpc.ClientConn)
}

func serveHTTP(l net.Listener, srv *grpc.Server, cfg *GRPCServerConfig) error {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	defer lis.Close()

	go srv.Serve(lis)

	grpc.EnableTracing = EnableTracing

	dialCtx := context.Background()
	dialCtx, cancel := context.WithTimeout(dialCtx, 10*time.Second)
	defer cancel()

	cc, err := grpc.DialContext(dialCtx, lis.Addr().String(),
		grpc.WithAuthority(l.Addr().String()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		return err
	}
	defer cc.Close()

	r := chi.NewRouter()
	r.Use(middleware.WithValue(ctxKeyClientConn, cc))
	r.Use(pubsubPushRequestConvertMiddleware)
	for _, desc := range cfg.UseHTTP.Descriptors {
		registerServiceDescHTTP(r, desc)
	}
	// debug begin
	r.Get("/debug/events", trace.Events)
	r.Get("/debug/requests", trace.Traces)
	if err := chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Debugf("[%s] %s", method, route)
		return nil
	}); err != nil {
		panic(err)
	}
	// debug end

	var hs http.Server
	hs.Handler = r
	return hs.Serve(l)
}

// PubSub Pushリクエストの変換処理
// 特記事項
// PubSubPushのHTTPリクエストペイロードには、歴史的経緯から messageId と message_id 、 publishTime と publish_time 両方のプロパティが含まれている。
// As-Isでprotojson.Unmarshalにかけると、フィールド重複エラーによりパースに失敗してしまう。
// 回避のため、重複フィールドを排除し、パイプラインへ流し直す。
// 本システムとしてx-authorizationヘッダとcookieヘッダをPubsub Publish時にAttributeとして、
// 格納し連携する仕組みのため、本関数でAttributeを解析し、HttpHeaderに追加する。
func pubsubPushRequestConvertMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := ctxlogrus.Extract(r.Context())

		// UserAgentベースでPubSubからの呼び出しか判定
		if !strings.HasPrefix(r.UserAgent(), "APIs-Google") {
			// Google APIからでなければ、そのまま流す
			next.ServeHTTP(w, r)
			return
		}

		rawBody, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		trimReq := struct {
			Message      *pubsub.PubsubMessage `json:"message"`
			Subscription string                `json:"subscription"`
		}{}

		// json.Unmarshalすることで「message_id」「publish_time」を除外
		if err := json.Unmarshal(rawBody, &trimReq); err != nil {
			// Unmarshalに失敗したら、そのまま流す
			logger.WithError(err).Error("failed to unmarshal")
			next.ServeHTTP(w, r)
			return
		}

		if trimReq.Subscription == "" || trimReq.Message == nil {
			// データが空なら、そのまま流す
			logger.Info("not pubsub push request")
			next.ServeHTTP(w, r)
			return
		}

		attrs := trimReq.Message.Attributes

		// attributeなどはheaderとして渡す
		r.Header.Add(http.CanonicalHeaderKey("x-pubsub-subscription"), trimReq.Subscription)
		if data, err := json.Marshal(attrs); err != nil {
			logger.WithError(err).Error("failed to marshal attribute.")
			panic("failed to marshal attribute.")
		} else {
			// valueに格納される値がHTTP Headerに使用可能な文字制限に抵触しないように、
			// Base64 Encodeして格納する
			base64Data := base64.StdEncoding.EncodeToString(data)
			r.Header.Add(http.CanonicalHeaderKey("x-pubsub-attributes"), base64Data)
		}

		// pubsubリクエストから認証情報を取得してメタデータに設定
		if v, ok := attrs["cookie"]; ok {
			var values []string
			if err := json.Unmarshal([]byte(v), &values); err != nil {
				logger.WithError(err).Error("failed to unmarshal cookie")
				// 異常時は追加を諦める
			} else {
				for _, value := range values {
					r.Header.Add("cookie", value)
				}
			}
		}

		if v, ok := attrs["x-authorization_00"]; ok {
			var values []string
			if err := json.Unmarshal([]byte(v), &values); err != nil {
				logger.WithError(err).Error("failed to unmarshal x-authorization")
				// 異常時は追加を諦める
			} else {
				for _, value := range values {
					r.Header.Add("x-authorization", value)
				}
			}
		}

		// dataフィールドはBase64 EncodeされているためDecode
		b64DecString, err := base64.StdEncoding.DecodeString(trimReq.Message.Data)
		if err != nil {
			logger.WithError(err).Error("failed to decode message data")
			// 異常時は諦めて元のまま流す
			next.ServeHTTP(w, r)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b64DecString))

		next.ServeHTTP(w, r)

	})
}

func registerServiceDescHTTP(r chi.Router, desc grpc.ServiceDesc) {
	rt := reflect.TypeOf(desc.HandlerType).Elem()
	r.Route("/"+desc.ServiceName, func(r chi.Router) {
		for _, method := range desc.Methods {
			rm, ok := rt.MethodByName(method.MethodName)
			if !ok {
				panic(fmt.Sprintf("unable to get method by name: type=%v methodName=%v", rt, method.MethodName))
			}

			argTyp := rm.Type.In(1).Elem()
			arg := reflect.New(argTyp).Interface()

			replyTyp := rm.Type.Out(0).Elem()
			reply := reflect.New(replyTyp).Interface()

			var hdlr methodInvokerHandler
			hdlr.FullName = "/" + desc.ServiceName + "/" + method.MethodName
			hdlr.Arg = arg.(proto.Message)
			hdlr.Reply = reply.(proto.Message)

			r.Post("/"+method.MethodName, hdlr.ServeHTTP)
		}
	})
}

type methodInvokerHandler struct {
	FullName string

	Arg proto.Message

	Reply proto.Message
}

var (
	// See https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
	codeMappings = map[codes.Code]int{
		codes.OK:                 http.StatusOK,
		codes.Canceled:           499, // Client Closed Request
		codes.Unknown:            http.StatusInternalServerError,
		codes.InvalidArgument:    http.StatusBadRequest,
		codes.DeadlineExceeded:   http.StatusGatewayTimeout,
		codes.NotFound:           http.StatusNotFound,
		codes.AlreadyExists:      http.StatusConflict,
		codes.PermissionDenied:   http.StatusForbidden,
		codes.ResourceExhausted:  http.StatusTooManyRequests,
		codes.FailedPrecondition: http.StatusBadRequest,
		codes.Aborted:            http.StatusConflict,
		codes.OutOfRange:         http.StatusBadRequest,
		codes.Unimplemented:      http.StatusNotImplemented,
		codes.Internal:           http.StatusInternalServerError,
		codes.Unavailable:        http.StatusServiceUnavailable,
		codes.DataLoss:           http.StatusInternalServerError,
		codes.Unauthenticated:    http.StatusUnauthorized,
	}
)

func (hdlr *methodInvokerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// See https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
	for name := range req.Header {
		values := req.Header.Values(name)
		switch strings.ToLower(name) {
		case "content-type":
			values = append(values[0:0], "application/grpc")
		case "content-length", "connection":
			values = values[0:0]
		}
		for _, value := range values {
			ctx = metadata.AppendToOutgoingContext(ctx, name, value)
		}
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	arg := proto.Clone(hdlr.Arg)
	if err := protojson.Unmarshal(reqBody, arg); err != nil {
		panic(err)
	}
	reply := proto.Clone(hdlr.Reply)

	var (
		metaHeader = metadata.MD{}
	)
	cc := clientConnFromContext(ctx)
	err = cc.Invoke(ctx, hdlr.FullName, arg, reply, grpc.Header(&metaHeader))
	rst := status.Convert(err)
	for k := range metaHeader {
		for _, v := range metaHeader.Get(k) {
			w.Header().Add(http.CanonicalHeaderKey(k), v)
		}
	}
	w.Header().Set("grpc-status", strconv.FormatUint(uint64(rst.Code()), 10))
	w.Header().Set("grpc-message", rst.Message())

	if v, ok := codeMappings[rst.Code()]; ok {
		w.WriteHeader(v)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if rst.Code() == codes.OK {
		resBody, err := protojson.Marshal(reply)
		if err != nil {
			panic(err)
		}
		if _, err := w.Write(resBody); err != nil {
			panic(err)
		}
	}
}
