package visibility

import (
	"fmt"
	"mtechnavi/sharelib/dataproxy/codec"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	dataproxypb "mtechnavi/sharelib/protobuf/record"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type Message interface {
	Type() string
	New() Message
	Set(f Field)
	Range(fn func(f Field) bool)
	Proto() proto.Message
}

type Field interface {
	New(v any) Field
	Name() string
	Value() any
	Zero() any
	Message() Message
	IsValid() bool
	Descriptor() Descriptor
}

type Descriptor interface {
	Scope() mtnpb.Scope
	IsMessage() bool
}

var (
	_ Message    = &protoMessage{}
	_ Message    = &recordMessage{}
	_ Field      = &protoField{}
	_ Field      = &recordField{}
	_ Descriptor = &protoDescriptor{}
	_ Descriptor = &recordDescriptor{}
)

func NewScopedMessage(scope mtnpb.Scope, msg proto.Message) proto.Message {
	src := newMessage(scope, msg)
	dst := newScopedMessage(scope, src)
	return dst.Proto()
}

func newMessage(scope mtnpb.Scope, msg proto.Message) Message {
	switch v := msg.(type) {
	case *dataproxypb.Record:
		return &recordMessage{v, scope}
	default:
		return &protoMessage{v.ProtoReflect(), scope}
	}
}

func newScopedMessage(scope mtnpb.Scope, m Message) Message {
	dst := m.New()

	pkName := codec.GetPrimaryKeyName(m.Type())
	m.Range(func(f Field) bool {
		for _, name := range []string{
			pkName,
			"created_at",
			"updated_at",
			"deleted_at",
			"shared_properties",
		} {
			if f.Name() == name {
				dst.Set(f)
				return true
			}
		}
		if fld := newScopedField(scope, f); fld != nil {
			dst.Set(fld)
		}
		return true
	})
	return dst
}

func newScopedField(scope mtnpb.Scope, f Field) Field {
	// ゼロ値の場合、空フィールドをそのまま返す
	if !f.IsValid() {
		return f
	}
	fd := f.Descriptor()
	if fd.Scope() != scope {
		// スコープ外の場合、ゼロ値で更新をかけることを考慮して空フィールドを返す
		return f.New(f.Zero())
	}
	var fld Field
	if fd.IsMessage() {
		fld = f.New(newScopedMessage(scope, f.Message()))
	} else {
		fld = f
	}
	return fld
}

func getScope(parent mtnpb.Scope, scope mtnpb.Scope) mtnpb.Scope {
	switch parent {
	case mtnpb.Scope_PUBLIC:
		return getPublicFieldScope(scope)
	case mtnpb.Scope_GRANTED:
		return getGrantedFieldScope(scope)
	case mtnpb.Scope_PRIVATE:
		return getPrivateFieldScope(scope)
	default:
		panic(fmt.Sprintf("unsupported parent scope: %v", parent))
	}
}

func getPublicFieldScope(s mtnpb.Scope) mtnpb.Scope {
	switch s {
	case mtnpb.Scope_INHERIT:
		return mtnpb.Scope_PUBLIC
	case mtnpb.Scope_PUBLIC:
		return mtnpb.Scope_PUBLIC
	case mtnpb.Scope_GRANTED:
		return mtnpb.Scope_GRANTED
	case mtnpb.Scope_PRIVATE:
		return mtnpb.Scope_PRIVATE
	default:
		panic(fmt.Sprintf("unsupported visibility scope: %v", s))
	}
}

func getGrantedFieldScope(s mtnpb.Scope) mtnpb.Scope {
	switch s {
	case mtnpb.Scope_INHERIT:
		return mtnpb.Scope_GRANTED
	case mtnpb.Scope_PUBLIC:
		return mtnpb.Scope_GRANTED
	case mtnpb.Scope_GRANTED:
		return mtnpb.Scope_GRANTED
	case mtnpb.Scope_PRIVATE:
		return mtnpb.Scope_PRIVATE
	default:
		panic(fmt.Sprintf("unsupported visibility scope: %v", s))
	}
}

func getPrivateFieldScope(s mtnpb.Scope) mtnpb.Scope {
	switch s {
	case mtnpb.Scope_INHERIT:
		return mtnpb.Scope_PRIVATE
	case mtnpb.Scope_PUBLIC:
		return mtnpb.Scope_PRIVATE
	case mtnpb.Scope_GRANTED:
		return mtnpb.Scope_PRIVATE
	case mtnpb.Scope_PRIVATE:
		return mtnpb.Scope_PRIVATE
	default:
		panic(fmt.Sprintf("unsupported visibility scope: %v", s))
	}
}

func isZeroValue(v interface{}) bool {
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
