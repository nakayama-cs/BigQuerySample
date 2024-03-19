package startup

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	"mtechnavi/sharelib/cache"

	log "github.com/sirupsen/logrus"
)

type Job interface {
	Init(ctx context.Context) error
	Execute(ctx context.Context) error
}

type JobConfig struct {
	Stdout io.Writer
	Stderr io.Writer
	Args   []string
	New    func(context.Context) (Job, error)
}

func RunJob(cfg *JobConfig) int {
	log.SetOutput(cfg.Stderr)
	var (
		loglevel string = "info"
		logfile  string = "-"
	)
	flgs := flag.NewFlagSet(filepath.Base(cfg.Args[0]), flag.ContinueOnError)
	flgs.StringVar(&loglevel, "loglevel", loglevel, "ログレベル (choices: trace, debug, info, warn, error)")
	flgs.StringVar(&logfile, "logfile", logfile, "ログ出力先ファイルパス")
	if err := flgs.Parse(cfg.Args[1:]); errors.Is(err, flag.ErrHelp) {
		return 1
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

	job, err := cfg.New(initCtx)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	if err := job.Execute(initCtx); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	return 0
}
