package visibility

import (
	"mtechnavi/sharelib/dataproxy/codec"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type protoMessage struct {
	m     protoreflect.Message
	scope mtnpb.Scope
}

func (m *protoMessage) Type() string {
	return string(proto.MessageName(m.Proto()))
}

func (m *protoMessage) New() Message {
	return &protoMessage{
		m:     m.m.New(),
		scope: m.scope,
	}
}

func (m *protoMessage) Set(f Field) {
	if v, ok := f.(*protoField); ok {
		if f.IsValid() {
			m.m.Set(v.fd, v.v)
		}
	} else {
		panic("protoMessage.Set: invalid field type")
	}
}

func (m *protoMessage) Range(fn func(f Field) bool) {
	fds := m.m.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		ex := getFieldExtension(fd)
		f := protoField{
			scope: getScope(m.scope, ex.Scope),
			fd:    fd,
			v:     m.m.Get(fd),
		}
		// 値がZeroValueでないフィールドのみ絞り込めれば良いため、Hasを使用する
		// 全フィールドをループさせる必要がある場合は、Hasを除く。
		if m.m.Has(fd) && !fn(&f) {
			return
		}
	}
}

func (m protoMessage) Proto() proto.Message {
	return m.m.Interface()
}

type protoField struct {
	scope mtnpb.Scope
	fd    protoreflect.FieldDescriptor
	v     protoreflect.Value
}

func (f *protoField) New(v any) Field {
	fld := &protoField{
		scope: f.scope,
		fd:    f.fd,
	}
	if m, ok := v.(*protoMessage); ok {
		fld.v = protoreflect.ValueOfMessage(m.Proto().ProtoReflect())
	} else {
		if m, ok := v.(proto.Message); ok {
			if !isZeroValue(m) {
				fld.v = protoreflect.ValueOf(m.ProtoReflect())
			} else {
				fld.v = protoreflect.ValueOf(f.v.Message().New())
			}
		} else {
			switch v.(type) {
			case interface{}:
				fld.v = protoreflect.ValueOf(nil)
			default:
				fld.v = protoreflect.ValueOf(v)
			}
		}
	}
	return fld
}

func (f *protoField) Name() string {
	return f.fd.TextName()
}

func (f *protoField) Value() any {
	return f.v.Interface()
}

func (f *protoField) Zero() any {
	return codec.GetProtoZeroValue(f.fd, f.v)
}

func (f *protoField) Message() Message {
	return newMessage(f.Descriptor().Scope(), f.v.Message().Interface())
}

func (f *protoField) IsValid() bool {
	if f.Descriptor().IsMessage() {
		return f.v.IsValid() && f.v.Message().IsValid()
	} else {
		return f.v.IsValid()
	}
}

func (f *protoField) Descriptor() Descriptor {
	return &protoDescriptor{
		fd:    f.fd,
		scope: f.scope,
	}
}

type protoDescriptor struct {
	fd    protoreflect.FieldDescriptor
	scope mtnpb.Scope
}

func (d *protoDescriptor) Scope() mtnpb.Scope {
	return d.scope
}

func (d *protoDescriptor) IsMessage() bool {
	if d.fd == nil {
		return false
	}
	if d.fd.Cardinality() == protoreflect.Repeated {
		return false
	}
	return d.fd.Kind() == protoreflect.MessageKind
}

func getFieldExtension(fd protoreflect.FieldDescriptor) *mtnpb.VisibilityOptions {
	out := &mtnpb.VisibilityOptions{}
	if fd == nil {
		return out
	}
	proto.RangeExtensions(fd.Options(), func(xt protoreflect.ExtensionType, ev interface{}) bool {
		v, ok := ev.(*mtnpb.VisibilityOptions)
		if !ok {
			return true
		}
		out = v
		return false
	})
	return out
}
