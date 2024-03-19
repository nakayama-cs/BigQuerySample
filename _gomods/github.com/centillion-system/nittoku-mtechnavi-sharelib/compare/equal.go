package compare

import (
	"bytes"
	"math"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func equalField(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	switch {
	case fd.IsList():
		return equalList(fd, x.List(), y.List())
	case fd.IsMap():
		return equalMap(fd, x.Map(), y.Map())
	default:
		return equalValue(fd, x, y)
	}
}

func equalMap(fd protoreflect.FieldDescriptor, x, y protoreflect.Map) bool {
	if x.Len() != y.Len() {
		return false
	}
	equal := true
	x.Range(func(k protoreflect.MapKey, vx protoreflect.Value) bool {
		vy := y.Get(k)
		equal = y.Has(k) && equalValue(fd.MapValue(), vx, vy)
		return equal
	})
	return equal
}

func equalList(fd protoreflect.FieldDescriptor, x, y protoreflect.List) bool {
	if x.Len() != y.Len() {
		return false
	}
	for i := x.Len() - 1; i >= 0; i-- {
		mx, my := x.Get(i), y.Get(i)
		switch fd.Kind() {
		case protoreflect.MessageKind:
			var s state
			return len(s.compareMessage(mx.Message(), my.Message())) == 0
		}
		if !equalValue(fd, x.Get(i), y.Get(i)) {
			return false
		}
	}
	return true
}

func equalValue(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return x.Bool() == y.Bool()
	case protoreflect.EnumKind:
		return x.Enum() == y.Enum()
	case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return x.Int() == y.Int()
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return x.Uint() == y.Uint()
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		fx := x.Float()
		fy := y.Float()
		if math.IsNaN(fx) || math.IsNaN(fy) {
			return math.IsNaN(fx) && math.IsNaN(fy)
		}
		return fx == fy
	case protoreflect.StringKind:
		return x.String() == y.String()
	case protoreflect.BytesKind:
		return bytes.Equal(x.Bytes(), y.Bytes())
	default:
		return x.Interface() == y.Interface()
	}
}
