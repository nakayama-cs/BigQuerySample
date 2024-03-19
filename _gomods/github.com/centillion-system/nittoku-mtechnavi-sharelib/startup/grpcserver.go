package startup

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"mtechnavi/sharelib/cache"
	"net"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Stdout io.Writer

	Stderr io.Writer

	Args []string

	New func(context.Context) (*grpc.Server, error)

	// gRPCサーバを、HTTPを通して呼び出す構成とする場合に設定する
	UseHTTP *HTTPBindingConfig
}

type HTTPBindingConfig struct {
	Descriptors []grpc.ServiceDesc
}

func RunGRPCServer(cfg *GRPCServerConfig) int {
	log.SetOutput(cfg.Stderr)

	var (
		listenAddr string = DefaultListenAddr()
		loglevel   string = "info"
		logfile    string = "-"
	)
	flgs := flag.NewFlagSet(filepath.Base(cfg.Args[0]), flag.ContinueOnError)
	flgs.StringVar(&listenAddr, "l", listenAddr, "起動するgRPC Serverの待ち受けアドレス")
	flgs.StringVar(&loglevel, "loglevel", loglevel, "ログレベル (choices: trace, debug, info, warn, error)")
	flgs.StringVar(&logfile, "logfile", logfile, "ログ出力先ファイルパス")
	if err := flgs.Parse(cfg.Args[1:]); errors.Is(err, flag.ErrHelp) {
		return 0
	} else if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 128
	}

	loggerCloseFn, err := initLogger(loglevel, logfile, cfg.Stderr)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	defer loggerCloseFn()

	initCtx := context.Background()
	initCtx, cancel := context.WithTimeout(initCtx, 10*time.Second)
	defer cancel()

	lru, err := initCache(initCtx)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	initCtx = cache.ContextWithCache(initCtx, lru)

	if err := initMessage(initCtx); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	if err := initTrace(initCtx); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	srv, err := cfg.New(initCtx)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	log.Debugf("listening on %s ...", listenAddr)
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	defer l.Close()

	var serve func(net.Listener, *grpc.Server, *GRPCServerConfig) error
	if cfg.UseHTTP != nil {
		serve = serveHTTP
	} else {
		serve = serveGRPC
	}

	if err := serve(l, srv, cfg); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	return 0
}

func serveGRPC(l net.Listener, srv *grpc.Server, cfg *GRPCServerConfig) error {
	return srv.Serve(l)
}
