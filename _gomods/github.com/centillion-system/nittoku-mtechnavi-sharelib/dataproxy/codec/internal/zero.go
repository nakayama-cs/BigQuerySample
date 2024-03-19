package internal

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func IsZeroValue(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Map:
		return rv.Len() == 0
	}
	return rv.IsZero()
}

var (
	zeroBool    bool
	zeroInt32   int32
	zeroInt64   int64
	zeroUint32  uint32
	zeroUint64  uint64
	zeroFloat32 float32
	zeroFloat64 float64
	zeroString  string
	zeroEnum    protoreflect.EnumNumber
	zeroBytes   = make([]byte, 0)
	zeroSlice   = make([]interface{}, 0)
	zeroMap     = make(map[string]interface{}, 0)
)

func GetProtoZeroValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) any {
	if fd.Cardinality() == protoreflect.Repeated {
		if fd.IsList() {
			return zeroSlice
		}
		if fd.IsMap() {
			return zeroMap
		}
	}
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return zeroBool
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return zeroInt32
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return zeroInt64
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return zeroUint32
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return zeroUint64
	case protoreflect.FloatKind:
		return zeroFloat32
	case protoreflect.DoubleKind:
		return zeroFloat64
	case protoreflect.StringKind:
		return zeroString
	case protoreflect.BytesKind:
		return zeroBytes
	case protoreflect.EnumKind:
		return zeroEnum
	case protoreflect.MessageKind:
		// oneofフィールドの場合は、使用しないためnil値を返す
		if fd.ContainingOneof() != nil {
			return nil
		}
		return v.Message().Type().Zero().Interface()
	default:
		panic(fmt.Sprintf("unsupported zero value: kind=%v", fd.Kind()))
	}
}

func GetZeroValue(v reflect.Value) any {
	switch v.Kind() {
	case reflect.Bool:
		return zeroBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return zeroInt64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return zeroUint64
	case reflect.Float32, reflect.Float64:
		return zeroFloat64
	case reflect.String:
		return zeroString
	case reflect.Slice:
		return zeroSlice
	case reflect.Map:
		return zeroMap
	case reflect.Struct:
		return reflect.Zero(v.Type()).Interface()
	case reflect.Invalid:
		return nil
	default:
		panic(fmt.Sprintf("unsupported zero value: kind=%v", v.Kind()))
	}
}
