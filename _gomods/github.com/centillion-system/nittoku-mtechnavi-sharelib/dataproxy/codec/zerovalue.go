package codec

import (
	codecinternal "mtechnavi/sharelib/dataproxy/codec/internal"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func GetProtoZeroValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) any {
	return codecinternal.GetProtoZeroValue(fd, v)
}

func GetZeroValue(v reflect.Value) any {
	return codecinternal.GetZeroValue(v)
}
