package validator

import (
	"context"
	sharelibpb "mtechnavi/sharelib/protobuf"
	pb "mtechnavi/sharelib/validator/protobuf"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func validate(req interface{}) error {
	if v, ok := req.(interface {
		ValidateAll() error
	}); ok {
		return v.ValidateAll()
	} else {
		return nil
	}
}

func toValidationError(req interface{}, err error) *pb.ValidationError {
	if err == nil {
		panic("err is nil")
	}
	trans := &translator{
		Req: req,
	}

	errs := unwrap(err)
	var out pb.ValidationError
	for _, err := range errs {
		var v pb.Descriptor
		v.MessageName = err.Error()
		if ve, ok := err.(validationError); ok {
			v.Field = trans.Field(ve.Field())
			v.MessageName, v.MessageArgs = trans.Reason(ve.Reason())
		}
		out.Descriptors = append(out.Descriptors, &v)
	}
	return &out
}

func unwrap(err error) []error {
	if e, ok := err.(interface {
		AllErrors() []error
	}); ok {
		return e.AllErrors()
	} else {
		return []error{err}
	}
}

func sendMetadata(ctx context.Context, ve *pb.ValidationError) error {
	data, err := protojson.Marshal(ve)
	if err != nil {
		return err
	}

	md := metadata.Pairs("x-validation-error", string(data))
	if err := grpc.SendHeader(ctx, md); err != nil {
		return err
	}
	return nil
}

type Validator interface {
	ValidateRequest(ctx context.Context, req any) error
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := ctxlogrus.Extract(ctx)

		if err := validate(req); err != nil {
			ve := toValidationError(req, err)
			if err := sendMetadata(ctx, ve); err != nil {
				logger.WithError(err).Warn("failed to send ValidationError headers")
			}
			return nil, status.Error(codes.InvalidArgument, sharelibpb.Error_INVALID_ARGUMENT.String())
		}

		if validator, ok := info.Server.(Validator); ok {
			if err := validator.ValidateRequest(ctx, req); err != nil {
				// errはValidateRequest側で制御
				return nil, err
			}
		}

		res, err := handler(ctx, req)
		return res, err
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, &recvWrapper{ss})
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if err := validate(m); err != nil {
		ctx := s.Context()
		ve := toValidationError(m, err)
		if err := sendMetadata(ctx, ve); err != nil {
			// warn and ignore error
			logger := ctxlogrus.Extract(ctx)
			logger.WithError(err).Warn("failed to send ValidationError headers")
		}
		return status.Error(codes.InvalidArgument, sharelibpb.Error_INVALID_ARGUMENT.String())
	}
	return nil
}
