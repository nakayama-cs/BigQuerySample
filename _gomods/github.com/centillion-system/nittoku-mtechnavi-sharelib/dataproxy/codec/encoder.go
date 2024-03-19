package codec

import (
	"encoding/json"
	"io"
	codecinternal "mtechnavi/sharelib/dataproxy/codec/internal"
	sharelibpb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/protobuf/mtn"
	pb "mtechnavi/sharelib/protobuf/record"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Deprecated: Use "mtechnavi/sharelib/dataproxy/transportcodec".Encode instead.
func Encode(src proto.Message) (*pb.Record, error) {
	if m, ok := src.(*pb.Record); ok {
		return m, nil
	} else {
		return encodeMessage(src.ProtoReflect())
	}
}

func encodeMessage(rv protoreflect.Message) (*pb.Record, error) {
	md := rv.Descriptor()
	var out pb.Record
	out.TypeName = string(md.FullName())
	pkName := GetPrimaryKeyName(string(md.FullName()))
	out.Fields = map[int32]*pb.Field{}
	out.Indexs = []*pb.Index{}
	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		if v, i, err := encodeField(rv, fd); err != nil {
			return nil, err
		} else if v != nil {
			out.Fields[int32(fd.Number())] = v
			switch fd.Name() {
			case protoreflect.Name(pkName):
				out.RecordId = rv.Get(fd).String()
			case "shared_properties":
				switch fd.Kind() {
				case protoreflect.MessageKind:
					if p, ok := rv.Get(fd).Message().Interface().(*sharelibpb.EmbeddedSharedProperties); ok {
						out.SharedProperties = p
					}
				}
			case "created_at":
				out.CreatedAt = rv.Get(fd).Int()
			case "updated_at":
				out.UpdatedAt = rv.Get(fd).Int()
			case "deleted_at":
				out.DeletedAt = rv.Get(fd).Int()
			}
			if i != nil {
				out.Indexs = append(out.Indexs, i)
			}
		} else {
			// omit zero value
			f, i, err := newField(fd, GetProtoZeroValue(fd, rv.Get(fd)))
			if err != nil {
				return nil, err
			}
			out.Fields[int32(fd.Number())] = f
			if i != nil {
				out.Indexs = append(out.Indexs, i)
			}
		}
	}
	// omit zero value
	if len(out.Fields) == 0 {
		out.Fields = nil
	}
	return &out, nil
}

type mapKey struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func encodeField(rv protoreflect.Message, fd protoreflect.FieldDescriptor) (*pb.Field, *pb.Index, error) {
	var value interface{}
	if fd.Cardinality() == protoreflect.Repeated {
		// MapおよびListの型はCardinalityがRepeatedとなる
		if fd.IsMap() {
			src := rv.Get(fd).Map()
			var l [][]interface{}
			src.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				l = append(l, []interface{}{k, v})
				return true
			})
			var m = map[string]interface{}{}
			for _, v := range l {
				key := v[0].(protoreflect.MapKey)
				val := v[1].(protoreflect.Value)
				k, err := encodeMapKey(fd, key)
				if err != nil {
					return nil, nil, err
				}
				v, err := encodeValue(fd, val)
				if err != nil {
					return nil, nil, err
				}
				m[k] = v
			}
			value = m
		} else {
			src := rv.Get(fd).List()
			l := make([]interface{}, src.Len())
			for i := 0; i < src.Len(); i++ {
				if v, err := encodeValue(fd, src.Get(i)); err != nil {
					return nil, nil, err
				} else {
					l[i] = v
				}
			}
			value = l
		}
	} else {
		if v, err := encodeValue(fd, rv.Get(fd)); err != nil {
			return nil, nil, err
		} else {
			value = v
		}
	}
	// omit zero value
	if codecinternal.IsZeroValue(value) {
		return nil, nil, nil
	} else if v, ok := value.(*pb.Record); ok && len(v.Fields) == 0 {
		return nil, nil, nil
	}
	return newField(fd, value)
}

func newField(fd protoreflect.FieldDescriptor, value interface{}) (*pb.Field, *pb.Index, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, nil, err
	}
	var out pb.Field
	out.Name = string(fd.Name())
	out.Value = data
	ext := getFieldExtension(fd)
	if ext.VisibilityOptions != nil {
		out.Visibility = ext.VisibilityOptions
	} else {
		out.Visibility = &mtn.VisibilityOptions{
			// デフォルトはPRIVATE
			Scope: mtn.Scope_PRIVATE,
		}
	}
	if ext.DataproxyOptions == nil {
		return &out, nil, nil
	}
	var index pb.Index
	index.Name = ext.DataproxyOptions.Index.Name
	index.Value = data
	return &out, &index, nil
}

func encodeValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) (interface{}, error) {
	kind := fd.Kind()
	if fd.IsMap() {
		kind = fd.MapValue().Kind()
	}
	switch kind {
	case protoreflect.BoolKind:
		return v.Bool(), nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return v.Int(), nil
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return v.Uint(), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return v.Float(), nil
	case protoreflect.StringKind:
		return v.String(), nil
	case protoreflect.BytesKind:
		return v.Bytes(), nil
	// case protoreflect.EnumKind:
	case protoreflect.MessageKind:
		if fd.ContainingOneof() != nil {
			// 指定されていないoneofフィールドは、使用しないためnilを返す
			if !v.IsValid() || !v.Message().IsValid() {
				return nil, nil
			}
		}
		return encodeMessage(v.Message())
	default:
		return v.Interface(), nil
	}
}

func encodeMapKey(fd protoreflect.FieldDescriptor, key protoreflect.MapKey) (string, error) {
	mapKey := &mapKey{
		Type:  fd.MapKey().Kind().String(),
		Value: key.Interface(),
	}
	b, err := json.Marshal(mapKey)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type fieldExtension struct {
	VisibilityOptions *mtn.VisibilityOptions
	DataproxyOptions  *mtn.DataproxyOptions
}

func getFieldExtension(fd protoreflect.FieldDescriptor) *fieldExtension {
	// デフォルト値
	var out fieldExtension
	// フィールドオプションがあれば取得
	proto.RangeExtensions(fd.Options(), func(xt protoreflect.ExtensionType, ev interface{}) bool {
		if v, ok := ev.(*mtn.VisibilityOptions); ok {
			out.VisibilityOptions = v
		}
		if v, ok := ev.(*mtn.DataproxyOptions); ok {
			out.DataproxyOptions = v
		}
		return true
	})
	return &out
}

func Debug(w io.Writer, v *pb.Record) error {
	pp := map[string]interface{}{}
	pp["type_name"] = v.TypeName
	fields := map[int32]map[string]interface{}{}
	for n, field := range v.Fields {
		f := map[string]interface{}{}
		f["name"] = field.Name
		f["value"] = string(field.Value)
		fields[n] = f
	}
	pp["fields"] = fields

	je := json.NewEncoder(w)
	je.SetIndent("", "  ")
	return je.Encode(pp)
}
