package startup

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"mtechnavi/sharelib/migrate"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type MigratorConfig struct {
	Stdout io.Writer

	Stderr io.Writer

	Args []string

	RegisterMessages []proto.Message

	New func(context.Context) (*migrate.Migrator, error)
}

func RunMigrator(cfg *MigratorConfig) int {
	logrus.SetOutput(cfg.Stderr)
	var (
		loglevel       string = "info"
		logfile        string = "-"
		projectID      string
		collectionPath string
		planFile       string
	)
	flgs := flag.NewFlagSet(filepath.Base(cfg.Args[0]), flag.ContinueOnError)
	flgs.StringVar(&loglevel, "loglevel", loglevel, "ログレベル (choices: trace, debug, info, warn, error)")
	flgs.StringVar(&logfile, "logfile", logfile, "ログ出力先ファイルパス")
	flgs.StringVar(&projectID, "project-id", projectID, "マイグレート対象のfirestoreが稼動しているプロジェクトID (mandatory for plan, apply)")
	flgs.StringVar(&collectionPath, "collection-path", collectionPath, "マイグレート対象のfirestoreコレクションパス (mandatory for plan )")
	flgs.StringVar(&planFile, "planfile", planFile, "planファイルパス (mandatory)")
	flgs.Usage = func() {
		p := func(format string, args ...interface{}) {
			fmt.Fprintf(cfg.Stderr, format, args...)
		}
		p("Usage:\n")
		p("  %s [options] {action}\n", flgs.Name())
		p("\n")
		p("Actions:\n")
		p("  show\n")
		p("    planで作成したデータファイルから、マイグレート内容を表示する。\n")
		p("  plan\n")
		p("    firestoreから指定コレクションのデータを全件読み込み、apply用のデータファイルをローカル上に作成する。\n")
		p("  apply\n")
		p("    planで作成したデータファイルから、実際にfirestoreに対して書き込みを実施する。\n")
		p("\n")
		p("Available Options:\n")
		flgs.PrintDefaults()
	}
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

	for _, msg := range cfg.RegisterMessages {
		migrate.RegisterMessage(msg)
	}

	initCtx := context.Background()
	initCtx, cancel := context.WithTimeout(initCtx, 10*time.Second)
	defer cancel()

	migrator, err := cfg.New(initCtx)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	var actionFn func(context.Context) error
	switch flgs.Arg(0) {
	case "show":
		showReq := migrate.ShowRequest{
			PlanFile: planFile,
		}
		if !showReq.IsValid() {
			flgs.Usage()
			return 128
		}
		actionFn = func(ctx context.Context) error {
			return migrator.Show(ctx, showReq)
		}
	case "plan":
		planReq := migrate.PlanRequest{
			ProjectID:      projectID,
			CollectionPath: collectionPath,
			PlanFile:       planFile,
		}
		if !planReq.IsValid() {
			flgs.Usage()
			return 128
		}
		actionFn = func(ctx context.Context) error {
			return migrator.Plan(ctx, planReq)
		}
	case "apply":
		applyReq := migrate.ApplyRequest{
			ProjectID: projectID,
			PlanFile:  planFile,
		}
		if !applyReq.IsValid() {
			flgs.Usage()
			return 128
		}
		actionFn = func(ctx context.Context) error {
			return migrator.Apply(ctx, applyReq)
		}
	default:
		flgs.Usage()
		return 128
	}

	ctx := context.Background()
	if err := actionFn(ctx); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	return 0
}
