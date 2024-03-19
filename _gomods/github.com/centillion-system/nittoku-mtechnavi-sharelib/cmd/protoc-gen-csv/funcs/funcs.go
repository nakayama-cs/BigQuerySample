package funcs

import (
	"fmt"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	"reflect"
	"strings"

	sharelibpb "mtechnavi/sharelib/protobuf"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	Map = map[string]any{
		"getCsvFileOptions": getCsvFileOptions,
		"getColumnSpecs":    getColumnSpecs,
	}
)

func getCsvFileOptions(m *protogen.Message) *mtnpb.CsvFileOptions {
	v, _ := proto.GetExtension(m.Desc.Options(), mtnpb.E_CsvDefaultFormat).(*mtnpb.CsvFileOptions)
	return v
}

func maxColumn(m map[int]columnSpec) int {
	n := 0
	for col := range m {
		if n < col {
			n = col
		}
	}
	return n
}

func ValidateMessage(m *protogen.Message) []string {
	var out []string
	// column指定の重複不可
	memo := map[int32]struct{}{}
	for _, fp := range getFields(m) {
		opt, ok := getFieldExtension(fp)
		if !ok {
			continue
		}
		if _, dup := memo[opt.Column]; dup {
			out = append(out, fmt.Sprintf("duplicated column: %d", opt.Column))
		}
		memo[opt.Column] = struct{}{}
	}
	return out
}

func getColumnSpecs(m *protogen.Message) []columnSpec {
	specsMap := map[int]columnSpec{}
	for _, fp := range getFields(m) {
		opt, ok := getFieldExtension(fp)
		if !ok {
			continue
		}
		if _, dup := specsMap[int(opt.Column)]; dup {
			panic(fmt.Sprintf("duplicated column: %d", opt.Column))
		}
		specsMap[int(opt.Column)] = columnSpec{
			Column:  int(opt.Column),
			Field:   fp,
			Header:  opt.Header,
			Default: opt.Default,
		}
	}
	maxCols := maxColumn(specsMap)

	specs := make([]columnSpec, maxCols)
	for col, spec := range specsMap {
		specs[col-1] = spec
	}
	for i := range specs {
		if specs[i].Column == 0 {
			specs[i].Column = i + 1
		}
	}
	return specs
}

type fieldPath []*protogen.Field

func getFields(m *protogen.Message) []fieldPath {
	var out []fieldPath
	for _, f := range m.Fields {
		if f.Message != nil {
			if isDatetime(f.Desc) {
				out = append(out, fieldPath{f})
				continue
			}
			if isDate(f.Desc) {
				out = append(out, fieldPath{f})
				continue
			}
			if isAmount(f.Desc) {
				out = append(out, fieldPath{f})
				continue
			}
			for _, nested := range getFields(f.Message) {
				out = append(out, append(fieldPath{f}, nested...))
			}
		} else {
			out = append(out, fieldPath{f})
		}
	}
	return out
}

func getFieldExtension(fp fieldPath) (*mtnpb.CsvColumnOptions, bool) {
	f := fp[len(fp)-1]
	return GetFieldExtension(f.Desc)
}

func GetFieldExtension(fd protoreflect.FieldDescriptor) (*mtnpb.CsvColumnOptions, bool) {
	var out *mtnpb.CsvColumnOptions
	proto.RangeExtensions(fd.Options(), func(xt protoreflect.ExtensionType, ev interface{}) bool {
		v, ok := ev.(*mtnpb.CsvColumnOptions)
		if !ok {
			return true
		}
		out = v
		return false
	})
	return out, out != nil
}

type columnSpec struct {
	Column int

	Field fieldPath

	Header string

	Default string
}

func (v *columnSpec) IsValid() bool {
	return len(v.Field) > 0
}

func (v *columnSpec) IsList() bool {
	return v.Field[len(v.Field)-1].Desc.IsList()
}

func (v *columnSpec) IsDatetime() bool {
	if v.Kind() != protoreflect.MessageKind {
		return false
	}
	return isDatetime(v.Field[len(v.Field)-1].Desc)
}

func isDatetime(fd protoreflect.FieldDescriptor) bool {
	dt := new(sharelibpb.Datetime).ProtoReflect()
	return fd.Message().FullName() == dt.Descriptor().FullName()
}

func (v *columnSpec) IsDate() bool {
	if v.Kind() != protoreflect.MessageKind {
		return false
	}
	return isDate(v.Field[len(v.Field)-1].Desc)
}

func isDate(fd protoreflect.FieldDescriptor) bool {
	dt := new(sharelibpb.Date).ProtoReflect()
	return fd.Message().FullName() == dt.Descriptor().FullName()
}

func (v *columnSpec) IsAmount() bool {
	if v.Kind() != protoreflect.MessageKind {
		return false
	}
	return isAmount(v.Field[len(v.Field)-1].Desc)
}

func isAmount(fd protoreflect.FieldDescriptor) bool {
	amount := new(sharelibpb.Amount).ProtoReflect()
	return fd.Message().FullName() == amount.Descriptor().FullName()
}

func (v *columnSpec) Kind() protoreflect.Kind {
	return v.Field[len(v.Field)-1].Desc.Kind()
}

func (v *columnSpec) Index() int {
	return v.Column - 1
}

func (v *columnSpec) FieldAccess(recv string) string {
	return recv + "." + v.JoinFieldNames(".")
}

func (v *columnSpec) FieldAccessList(recv string) []string {
	out := make([]string, 0, len(v.Field)-1)
	for i := 0; i < len(v.Field)-1; i++ {
		out = append(out, recv+"."+v.joinFieldNames(recv, i+1))
	}
	return out
}

func (v *columnSpec) JoinFieldNames(sep string) string {
	return v.joinFieldNames(sep, len(v.Field))
}

func (v *columnSpec) joinFieldNames(sep string, n int) string {
	l := make([]string, n)
	for i, f := range v.Field[:n] {
		l[i] = f.GoName
	}
	return strings.Join(l, sep)
}

func (v *columnSpec) MessageName(i int) string {
	f := v.Field[i]
	if k := f.Desc.Kind(); k != protoreflect.MessageKind {
		panic(fmt.Sprintf("unsupported type: %v", k))
	}
	return f.Message.GoIdent.GoName
}

func (v *columnSpec) GoType() reflect.Type {
	if v.IsDatetime() {
		return reflect.TypeOf(new(sharelibpb.Datetime))
	}
	if v.IsDate() {
		return reflect.TypeOf(new(sharelibpb.Date))
	}
	if v.IsAmount() {
		return reflect.TypeOf(new(sharelibpb.Amount))
	}
	k := v.Field[len(v.Field)-1].Desc.Kind()
	typ, ok := map[protoreflect.Kind]reflect.Type{
		protoreflect.BoolKind:     reflect.TypeOf(false),
		protoreflect.Int32Kind:    reflect.TypeOf(int32(0)),
		protoreflect.Sint32Kind:   reflect.TypeOf(int32(0)),
		protoreflect.Uint32Kind:   reflect.TypeOf(uint32(0)),
		protoreflect.Int64Kind:    reflect.TypeOf(int64(0)),
		protoreflect.Sint64Kind:   reflect.TypeOf(int64(0)),
		protoreflect.Uint64Kind:   reflect.TypeOf(uint64(0)),
		protoreflect.Sfixed32Kind: reflect.TypeOf(int32(0)),
		protoreflect.Fixed32Kind:  reflect.TypeOf(uint32(0)),
		protoreflect.FloatKind:    reflect.TypeOf(float32(0)),
		protoreflect.Sfixed64Kind: reflect.TypeOf(int64(0)),
		protoreflect.Fixed64Kind:  reflect.TypeOf(uint64(0)),
		protoreflect.DoubleKind:   reflect.TypeOf(float64(0)),
		protoreflect.StringKind:   reflect.TypeOf(""),
	}[k]
	if !ok {
		panic(fmt.Sprintf("unsupported type: %v", k))
	}
	return typ
}
