package acl

import (
	pb "mtechnavi/idp/protobuf"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Masker struct {
	AccessControls *pb.SchemaAccessControl
}

func (m *Masker) Apply(v proto.Message) {
	m.apply(v.ProtoReflect())
}

func (m *Masker) apply(rv protoreflect.Message) {
	msgDesc := rv.Type().Descriptor()
	// 値がZeroValueかに関わらず動作させるため、Rangeでなく全フィールドをループさせる
	// マスク = ZeroValueで良いなら、単にRangeで良い
	fieldDescs := msgDesc.Fields()
	for i := 0; i < fieldDescs.Len(); i++ {
		fieldDesc := fieldDescs.Get(i)
		if !m.isAllowed(msgDesc, fieldDesc) {
			if rv.IsValid() {
				rv.Clear(fieldDesc)
			}
			continue
		}
		// repeatedやmapも、内包型がmessageであればmessageとなる
		if fieldDesc.Kind() != protoreflect.MessageKind {
			continue
		}
		switch {
		case fieldDesc.IsList():
			rl := rv.Get(fieldDesc).List()
			for i := 0; i < rl.Len(); i++ {
				m.apply(rl.Get(i).Message())
			}
		case fieldDesc.IsMap():
			rv.Get(fieldDesc).Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				m.apply(v.Message())
				return true
			})
		default:
			m.apply(rv.Get(fieldDesc).Message())
		}
	}
}

func (m *Masker) isAllowed(msgDesc protoreflect.MessageDescriptor, fieldDesc protoreflect.FieldDescriptor) bool {
	if isPrimaryKey(msgDesc, fieldDesc) {
		return true
	}
	propACL := m.findPropertyAccessControls(msgDesc)
	var (
		allowed bool
		name    = string(fieldDesc.Name())
	)
	for _, pat := range propACL.Allow {
		if pat == name || pat == "*" {
			allowed = true
			break
		}
	}
	if !allowed {
		return false
	}
	for _, pat := range propACL.Deny {
		if pat == name || pat == "*" {
			return false
		}
	}
	return true
}

var (
	snakeCaseReplacer = strings.NewReplacer(
		"A", "_a",
		"B", "_b",
		"C", "_c",
		"D", "_d",
		"E", "_e",
		"F", "_f",
		"G", "_g",
		"H", "_h",
		"I", "_i",
		"J", "_j",
		"K", "_k",
		"L", "_l",
		"M", "_m",
		"N", "_n",
		"O", "_o",
		"P", "_p",
		"Q", "_q",
		"R", "_r",
		"S", "_s",
		"T", "_t",
		"U", "_u",
		"V", "_v",
		"W", "_w",
		"X", "_x",
		"Y", "_y",
		"Z", "_z",
	)
)

func isPrimaryKey(msgDesc protoreflect.MessageDescriptor, fieldDesc protoreflect.FieldDescriptor) bool {
	pkName := snakeCaseReplacer.Replace(string(msgDesc.Name())) + "_id"
	pkName = strings.TrimPrefix(pkName, "_")
	return pkName == string(fieldDesc.Name())
}

func (m *Masker) findPropertyAccessControls(msgDesc protoreflect.MessageDescriptor) *pb.PropertyAccessControl {
	fullName := string(msgDesc.FullName())
	// wildcardよりも、厳密な型名でのマッチを優先させる
	for pat, v := range m.AccessControls.AccessControls {
		if pat == fullName {
			return v
		}
	}
	if v, ok := m.AccessControls.AccessControls["*"]; ok {
		return v
	}
	return &pb.PropertyAccessControl{}
}

func mergeAccessControls(l ...*pb.SchemaAccessControl) *pb.SchemaAccessControl {
	if len(l) == 0 {
		return &pb.SchemaAccessControl{}
	}

	var out pb.SchemaAccessControl
	for i := range l {
		for name, src := range l[i].AccessControls {
			src = proto.Clone(src).(*pb.PropertyAccessControl)
			if out.AccessControls == nil {
				out.AccessControls = map[string]*pb.PropertyAccessControl{}
			}
			if curr, ok := out.AccessControls[name]; ok {
				curr.Allow = append(curr.Allow, src.Allow...)
				curr.Deny = append(curr.Deny, src.Deny...)
				out.AccessControls[name] = curr
			} else {
				out.AccessControls[name] = src
			}
		}
	}
	// 不要な指定を削除し、結果をstableにするためにソートする
	for name, v := range out.AccessControls {
		v.Allow = cleanPropertyNames(v.Allow)
		v.Deny = cleanPropertyNames(v.Deny)
		out.AccessControls[name] = v
	}
	return &out
}

func cleanPropertyNames(l []string) []string {
	for _, s := range l {
		if s == "*" {
			return []string{"*"}
		}
	}
	sort.Strings(l)
	return l
}
