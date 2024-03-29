package bq

import (
	"errors"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var Done = errors.New("done")

type ConstraintProtoMessage[X any] interface {
	protoreflect.ProtoMessage
	*X
}
