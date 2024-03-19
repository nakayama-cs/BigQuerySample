package status

import (
	"context"
	"fmt"
	"mtechnavi/sharelib/constants"

	sharelibpb "mtechnavi/sharelib/protobuf"
	validatorpb "mtechnavi/sharelib/validator/protobuf"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// フィールドなし
func Error(ctx context.Context, code codes.Code, errName constants.MessageName, args ...any) error {
	return ErrorWithField(ctx, code, errName, "", args...)
}

// フィールドあり
func ErrorWithField(ctx context.Context, code codes.Code, errName constants.MessageName, field string, args ...any) error {
	logger := ctxlogrus.Extract(ctx)

	argsStr := make([]string, len(args))
	for i := range args {
		argsStr[i] = fmt.Sprint(args[i])
	}
	md, _ := metadata.FromOutgoingContext(ctx)
	md.Append("x-message-arguments", argsStr...)

	if err := grpc.SetHeader(ctx, md); err != nil {
		logger.Warnf("failed to set x-message-arguments")
	}

	descriptor := &validatorpb.Descriptor{
		Field:       field,
		MessageName: string(errName),
		MessageArgs: argsStr,
	}
	switch code {
	case codes.InvalidArgument:
		if err := setValidationErrorHeader(ctx, descriptor); err != nil {
			logger.WithError(err).Warnf("failed to set x-validation-error")
		}
	case codes.FailedPrecondition:
		if err := setValidationErrorHeader(ctx, descriptor); err != nil {
			logger.WithError(err).Warnf("failed to set x-validation-error")
		}
	default:
		// そのまま流す
	}
	return status.Error(code, string(errName))
}

func InvalidArgument(ctx context.Context, descriptor ...*validatorpb.Descriptor) *status.Status {
	logger := ctxlogrus.Extract(ctx)

	if err := setValidationErrorHeader(ctx, descriptor...); err != nil {
		logger.WithError(err).Error("failed to mershal ValidationError")
		return status.New(codes.Internal, sharelibpb.Error_INTERNAL_ERROR.String())
	}

	st := status.New(codes.InvalidArgument, sharelibpb.Error_INVALID_ARGUMENT.String())
	return st
}

func setValidationErrorHeader(ctx context.Context, descriptor ...*validatorpb.Descriptor) error {
	logger := ctxlogrus.Extract(ctx)

	ve := &validatorpb.ValidationError{
		Descriptors: descriptor,
	}

	data, err := protojson.Marshal(ve)
	if err != nil {
		logger.WithError(err).Error("failed to mershal ValidationError")
		return err
	}

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		logger.Error("failed to get outgoing context")
		return err
	}
	md.Append("x-validation-error", string(data))
	if err := grpc.SetHeader(ctx, md); err != nil {
		logger.WithError(err).Error("failed to send header")
		return err
	}

	return nil
}
