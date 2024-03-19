package encoding

import "google.golang.org/protobuf/reflect/protoreflect"

var (
	// for memo
	_ = []protoreflect.Kind{
		protoreflect.BoolKind,
		protoreflect.EnumKind,
		protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.FloatKind,
		protoreflect.Sfixed64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.DoubleKind,
		protoreflect.StringKind,
		protoreflect.BytesKind,
		protoreflect.MessageKind,
		protoreflect.GroupKind,
	}
)

type TypeEncoder interface {
	EncodeString(src any) (any, error)
	EncodeInt(src any) (any, error)
	EncodeUint(src any) (any, error)
	EncodeBytes(src any) (any, error)
	EncodeBool(src any) (any, error)
	EncodeFloat(src any) (any, error)
	EncodeEnum(src any) (any, error)
	EncodeMessage(src any) (any, error)
	EncodeGroup(src any) (any, error)
}

type TypeDecoder interface {
}

func oneof(v any, l ...any) bool {
	for _, val := range l {
		if val == v {
			return true
		}
	}
	return false
}

func isString(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.StringKind)
}

func isInt(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Fixed64Kind)
}

func isUint(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind)
}

func isBytes(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.BytesKind)
}

func isBool(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.BoolKind)
}

func isFloat(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.FloatKind,
		protoreflect.DoubleKind)
}

func isEnum(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.EnumKind)
}

func isMessage(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.MessageKind)
}

func isGroup(k protoreflect.Kind) bool {
	return oneof(k,
		protoreflect.GroupKind)
}

func EncodeType(enc TypeEncoder, src any) (any, error) {
	return nil, nil
}

func DecodeType(dec TypeDecoder, src any) (any, error) {
	return nil, nil
}
