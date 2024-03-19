package visibility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mtechnavi/sharelib/dataproxy/codec"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	dataproxypb "mtechnavi/sharelib/protobuf/record"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type recordMessage struct {
	m     *dataproxypb.Record
	scope mtnpb.Scope
}

func (m *recordMessage) Type() string {
	return m.m.TypeName
}

func (m *recordMessage) New() Message {
	n := &recordMessage{
		m: &dataproxypb.Record{
			RecordId:         m.m.RecordId,
			TypeName:         m.m.TypeName,
			Fields:           map[int32]*dataproxypb.Field{},
			CreatedAt:        m.m.CreatedAt,
			UpdatedAt:        m.m.UpdatedAt,
			DeletedAt:        m.m.DeletedAt,
			Indexs:           m.m.Indexs,
			SharedProperties: m.m.SharedProperties,
		},
		scope: m.scope,
	}
	if protoreflect.ValueOf(m.m.SharedProperties.ProtoReflect()).IsValid() {
		n.m.SharedProperties = m.m.SharedProperties
	}
	return n
}

func (m *recordMessage) Set(f Field) {
	if v, ok := f.(*recordField); ok {
		m.m.Fields[v.index] = v.f
	} else {
		panic(fmt.Sprintf("recordMessage.Set: unsupported field type=%T", f))
	}
}

func (m *recordMessage) Range(fn func(f Field) bool) {
	for idx, field := range m.m.Fields {
		var v interface{}
		if err := json.Unmarshal(field.Value, &v); err != nil {
			panic(err)
		}
		var fieldScope mtnpb.Scope
		if field.Visibility != nil {
			fieldScope = field.Visibility.Scope
		}
		f := recordField{
			scope: getScope(m.scope, fieldScope),
			index: idx,
			f:     field,
			v:     v,
		}
		if !fn(&f) {
			return
		}
	}
}

func (m recordMessage) Proto() proto.Message {
	return m.m
}

type recordField struct {
	scope mtnpb.Scope
	index int32
	f     *dataproxypb.Field
	v     interface{}
}

func (f *recordField) New(v any) Field {
	field := &recordField{
		scope: f.scope,
		index: f.index,
		f:     f.f,
	}
	var value interface{}
	if m, ok := v.(*recordMessage); ok {
		value = m.m
	} else {
		value = v
	}
	b, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	field.f.Value = b
	return field
}

func (f *recordField) Name() string {
	return f.f.Name
}

func (f *recordField) Value() any {
	return f.f.Value
}

func (f *recordField) Zero() any {
	if f.Descriptor().IsMessage() {
		return zeroRecord(f.Message())
	} else {
		return codec.GetZeroValue(reflect.ValueOf(f.v))
	}
}

func (f *recordField) Message() Message {
	var r dataproxypb.Record
	jd := json.NewDecoder(bytes.NewReader(f.f.Value))
	jd.UseNumber()
	if err := jd.Decode(&r); err != nil {
		panic(err)
	}
	return newMessage(f.Descriptor().Scope(), &r)
}

func (f *recordField) IsValid() bool {
	if f.Descriptor().IsMessage() {
		return !isZeroRecord(f.Message())
	} else {
		return !isZeroValue(f.v)
	}
}

func (f *recordField) Descriptor() Descriptor {
	return &recordDescriptor{
		f:     f.f,
		v:     f.v,
		scope: f.scope,
	}
}

type recordDescriptor struct {
	f     *dataproxypb.Field
	v     interface{}
	scope mtnpb.Scope
}

func (d *recordDescriptor) Scope() mtnpb.Scope {
	return d.scope
}

func (d *recordDescriptor) IsMessage() bool {
	return isRecord(d.v)
}

func isRecord(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return false
	}
	for _, k := range []string{
		// "record_id",
		"type_name",
		"fields",
	} {
		mv := rv.MapIndex(reflect.ValueOf(k))
		if !mv.IsValid() {
			return false
		}
	}
	return true
}

func isZeroRecord(v Message) bool {
	if v, ok := v.(*recordMessage); !ok {
		panic(fmt.Sprintf("isZeroRecord: unsupported Message type=%T", v))
	} else {
		var notEmpty bool
		v.Range(func(f Field) bool {
			fv := f.(*recordField)
			if fv.Descriptor().IsMessage() {
				if !isZeroRecord(fv.Message()) {
					notEmpty = true
				}
			} else {
				if !isZeroValue(fv.v) {
					notEmpty = true
				}
			}
			return true
		})
		return !notEmpty
	}
}

func zeroRecord(v Message) *dataproxypb.Record {
	if v, ok := v.(*recordMessage); !ok {
		panic(fmt.Sprintf("zeroRecord: unsupported Message type=%T", v))
	} else {
		var r dataproxypb.Record
		r.TypeName = v.m.TypeName
		r.Fields = map[int32]*dataproxypb.Field{}
		v.Range(func(f Field) bool {
			fv := f.(*recordField)
			var value interface{}
			if fv.Descriptor().IsMessage() {
				value = zeroRecord(fv.Message())
			} else {
				value = codec.GetZeroValue(reflect.ValueOf(fv.v))
			}
			b, err := json.Marshal(value)
			if err != nil {
				panic(fmt.Sprintf("unknown zero value: %v", err))
			}
			r.Fields[fv.index] = &dataproxypb.Field{
				Name:       fv.f.Name,
				Value:      b,
				Visibility: fv.f.Visibility,
			}
			return true
		})
		return &r
	}
}
