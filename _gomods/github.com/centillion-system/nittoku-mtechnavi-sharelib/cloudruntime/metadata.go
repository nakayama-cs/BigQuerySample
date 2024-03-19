package cloudruntime

import (
	"context"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

func ProjectID(ctx context.Context) string {
	logger := ctxlogrus.Extract(ctx)

	v := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if v != "" {
		return v
	}
	logger.Info("getting project-id from metadata ...")
	v, err := metadata.ProjectID()
	if err != nil {
		logger.WithError(err).Panic("unable to get projectId from metadata")
	}
	return v
}
