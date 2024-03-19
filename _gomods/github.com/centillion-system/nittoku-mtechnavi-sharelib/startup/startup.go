package startup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mtechnavi/sharelib/text/message"
	"net"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/storage"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"golang.org/x/text/language"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func DefaultListenAddr() string {
	// Cloud Runでの動作を前提とし、環境変数PORTで与えられたポート番号で待ち受ける
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8000"
	}
	return ":" + portStr
}

func CanonicalListenAddr(addr string) string {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		panic(err)
	}
	if host == "" {
		host = "localhost"
	}
	return net.JoinHostPort(host, portStr)
}

func initLogger(loglevel, logfile string, stderr io.Writer) (func(), error) {
	if lvl, err := logrus.ParseLevel(loglevel); err != nil {
		return nil, err
	} else {
		logrus.SetLevel(lvl)
	}

	// CloudRun上ではJSONログ、ローカルではテキストログ
	if metadata.OnGCE() {
		// See https://cloud.google.com/logging/docs/structured-logging#special-payload-fields
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyMsg:   "message",
			},
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
			FullTimestamp:   true,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyMsg:   "message",
			},
		})
	}

	var closer io.Closer
	if logfile == "-" || logfile == "" {
		logrus.SetOutput(stderr)
	} else {
		wc, err := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		closer = wc
		logrus.SetOutput(wc)
	}
	return func() {
		if closer == nil {
			return
		}
		if err := closer.Close(); err != nil {
			logrus.WithError(err).Errorf("failed to close a file: %q", logfile)
		}
	}, nil
}

func initMessage(ctx context.Context) error {
	logger := ctxlogrus.Extract(ctx)

	// TODO:将来的にHTTPアクセスに変更されるため、環境変数名はこのままとする
	const envName = "MTECHNAVI_MESSAGE_DICTIONARY_BASE_URL"
	baseURLStr := os.Getenv(envName)
	if baseURLStr == "" {
		logger.Warnf("environment variable %q is not set, skip to initialize message.", envName)
		return nil
	}
	message.Loader = func(ctx context.Context, lang language.Tag) (map[string]string, error) {
		base, _, _ := lang.Raw()
		// ja or en, etc
		langName := base.String()

		client, err := storage.NewClient(ctx)
		if err != nil {
			logger.WithError(err).Error("failed to init GCS Client")
			return nil, err
		}
		defer client.Close()

		rc, err := client.Bucket(baseURLStr).Object("language/" + langName + ".json").NewReader(ctx)
		if err != nil {
			logger.WithError(err).Error("failed to init GCS reader")
			return nil, err
		}
		defer rc.Close()

		var data map[string]string
		jd := json.NewDecoder(rc)
		if err := jd.Decode(&data); err != nil {
			return nil, err
		}
		return data, nil
	}

	// TODO:将来的にHTTPアクセスに変更されるため、コードは残す
	// baseURL, err := url.Parse(baseURLStr)
	// if err != nil {
	// 	return fmt.Errorf("invalid environment variable %q: %v", envName, err)
	// }
	// message.Loader = func(ctx context.Context, lang language.Tag) (map[string]string, error) {
	// 	base, _, _ := lang.Raw()
	// 	// ja or en, etc
	// 	langName := base.String()

	// 	u := *baseURL
	// 	u.Path = path.Join(u.Path, langName+".json")
	// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	// 	if err != nil {
	// 		logger.Panic(err)
	// 	}

	// 	hc, _, err := htransport.NewClient(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	res, err := hc.Do(req)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer res.Body.Close()

	// 	if res.StatusCode != http.StatusOK {
	// 		b, err := httputil.DumpResponse(res, true)
	// 		if err != nil {
	// 			logger.Panicf("go response %v: %v", res.Status, err)
	// 		}
	// 		return nil, &url.Error{
	// 			Op:  req.Method,
	// 			URL: req.URL.String(),
	// 			Err: errors.New(string(b)),
	// 		}
	// 	}

	// 	var data map[string]string
	// 	jd := json.NewDecoder(res.Body)
	// 	if err := jd.Decode(&data); err != nil {
	// 		return nil, err
	// 	}
	// 	return data, nil
	// }
	return nil
}

func initTrace(ctx context.Context) error {
	logger := ctxlogrus.Extract(ctx)

	const envExportProjectName = "MTECHNAVI_TRACE_EXPORT_PROJECT"
	exportProjectId := os.Getenv(envExportProjectName)
	if exportProjectId == "" {
		mes := fmt.Sprintf("environment variable %q is not set.", envExportProjectName)
		logger.Info(mes)
		return errors.New(mes)
	}

	traceFraction := 0.0
	const envFractionName = "MTECHNAVI_TRACE_SAMPLING_RATE"
	traceFractionStr := os.Getenv(envFractionName)
	if traceFractionStr == "" {
		logger.Infof("environment variable %q is not set,use default value.(tracing doesn't work)", envFractionName)
	} else {
		if v, err := strconv.ParseFloat(traceFractionStr, 64); err != nil {
			logger.WithError(err).Errorf("invalid environment variable %q. value:%s", envFractionName, traceFractionStr)
			return err
		} else {
			traceFraction = v
		}
	}

	exporter, err := texporter.New(texporter.WithProjectID(exportProjectId))
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(traceFraction)),
	)

	otel.SetTracerProvider(tp)

	return nil
}
