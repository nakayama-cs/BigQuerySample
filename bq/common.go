package bq

import (
	"errors"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var Done = errors.New("done")
var ErrCast = errors.New("cast error")

type ConstraintProtoMessage[X any] interface {
	protoreflect.ProtoMessage
	*X
}
