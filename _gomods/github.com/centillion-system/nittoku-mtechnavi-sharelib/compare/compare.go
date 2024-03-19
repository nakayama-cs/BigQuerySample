package compare

import (
	"errors"
	"fmt"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	"mtechnavi/sharelib/visibility"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type state struct {
	paths []pathValue
	// ignorePaths []string
}

type pathValue struct {
	path   string // json name path
	fd     protoreflect.FieldDescriptor
	valueX protoreflect.Value
	valueY protoreflect.Value
}

func (p *pathValue) values() (vx, vy any) {
	return toInterfaces(p.fd, p.valueX, p.valueY)
}

type Diff struct {
	Path   string // json name path
	ValueX any
	ValueY any
}

func Equal(scope mtnpb.Scope, x, y proto.Message) ([]*Diff, error) {
	if x == nil || y == nil {
		return nil, errors.New("comapared values should not be nil")
	}
	if reflect.TypeOf(x).Kind() == reflect.Ptr && x == y {
		return []*Diff{}, nil
	}
	rvx := visibility.NewScopedMessage(scope, x).ProtoReflect()
	rvy := visibility.NewScopedMessage(scope, y).ProtoReflect()
	// Descriptor()は型情報のみ返すので、Descriptorで型比較を行う
	if rvx.Descriptor() != rvy.Descriptor() {
		return nil, fmt.Errorf("type mismatch: x=%T, y=%T", x, y)
	}
	var s state
	return s.compareMessage(rvx, rvy), nil
}

func (s *state) compareMessage(x, y protoreflect.Message) (diffs []*Diff) {
	s.setPathValues(x, y)
	for _, pv := range s.paths {
		if !equalField(pv.fd, pv.valueX, pv.valueY) {
			var d Diff
			d.Path = pv.path
			d.ValueX, d.ValueY = pv.values()
			diffs = append(diffs, &d)
		}
	}
	return diffs
}

func (s *state) setPathValues(x, y protoreflect.Message) {
	fds := x.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
		mfd := fds.Get(i)
		if !x.Has(mfd) && !y.Has(mfd) {
			continue
		}
		var base pathValue
		s.anyValues(mfd, &base, x.Get(mfd), y.Get(mfd))
	}
}

func (s *state) anyValues(fd protoreflect.FieldDescriptor, v *pathValue, x, y protoreflect.Value) {
	switch fd.Kind() {
	case protoreflect.MessageKind:
		s.messageValues(fd, v, x, y)
	default:
		s.appendPath(fd, v, x, y)
	}
}

func (s *state) messageValues(fd protoreflect.FieldDescriptor, base *pathValue, x, y protoreflect.Value) {
	if fd.IsList() || fd.IsMap() {
		s.appendPath(fd, base, x, y)
		return
	}
	mx := x.Message()
	my := y.Message()
	fds := mx.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
		mfd := fds.Get(i)
		if !mx.Has(mfd) && !my.Has(mfd) {
			continue
		}
		fv := pathValue{
			path: strings.Join([]string{base.path, string(fd.JSONName())}, "."),
			fd:   mfd,
		}
		s.anyValues(mfd, &fv, mx.Get(mfd), my.Get(mfd))
	}
}

func (s *state) appendPath(fd protoreflect.FieldDescriptor, v *pathValue, x, y protoreflect.Value) {
	val := pathValue{
		path: strings.Join([]string{v.path, fd.JSONName()}, "."),
		fd:   fd,
	}
	val.path = strings.TrimPrefix(val.path, ".")
	val.valueX = x
	val.valueY = y
	s.paths = append(s.paths, val)
}

func toInterfaces(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) (vx, vy any) {
	switch {
	case fd.IsList():
		return toList(fd, x.List()), toList(fd, y.List())
	case fd.IsMap():
		return toMap(fd, x.Map()), toMap(fd, y.Map())
	case fd.Kind() == protoreflect.MessageKind:
		return x.Message().Interface(), y.Message().Interface()
	default:
		return x.Interface(), y.Interface()
	}
}

func toList(fd protoreflect.FieldDescriptor, list protoreflect.List) any {
	if list.Len() < 1 {
		return nil
	}
	var l reflect.Value
	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)

		var v any
		switch fd.Kind() {
		case protoreflect.MessageKind:
			v = item.Message().Interface()
		default:
			v = item.Interface()
		}
		if !l.IsValid() || l.IsNil() {
			l = makeSlice(v)
		}
		l = reflect.Append(l, reflect.ValueOf(v))
	}
	return l.Interface()
}

func toMap(fd protoreflect.FieldDescriptor, rm protoreflect.Map) any {
	var m reflect.Value
	if rm.Len() == 0 {
		var v map[string]any
		return v
	}
	rm.Range(func(mk protoreflect.MapKey, mv protoreflect.Value) bool {
		var k any
		var v any
		k = mk.Interface()
		switch fd.MapValue().Kind() {
		case protoreflect.MessageKind:
			// MapのValue値として、sliceやmapを指定することは出来ないためMessageKindの場合はProtoMessageとなる
			v = mv.Message().Interface()
		default:
			v = mv.Interface()
		}
		if !m.IsValid() || m.IsNil() {
			m = makeMap(k, v)
		}
		rk := reflect.ValueOf(k)
		rv := reflect.ValueOf(v)
		m.SetMapIndex(rk, rv)
		return true
	})
	return m.Interface()
}

func makeSlice(v any) reflect.Value {
	t := reflect.ValueOf(v).Type()
	return reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
}

func makeMap(k, v any) reflect.Value {
	tk := reflect.ValueOf(k).Type()
	tv := reflect.ValueOf(v).Type()
	return reflect.MakeMap(reflect.MapOf(tk, tv))
}
