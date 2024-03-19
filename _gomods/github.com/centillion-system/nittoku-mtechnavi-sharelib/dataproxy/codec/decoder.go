package codec

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	codecinternal "mtechnavi/sharelib/dataproxy/codec/internal"
	pb "mtechnavi/sharelib/protobuf/record"
	"reflect"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Deprecated: Use "mtechnavi/sharelib/dataproxy/transportcodec".Decode instead.
func Decode(dst proto.Message, data *pb.Record) error {
	switch dst.(type) {
	case *pb.Record:
		dist := reflect.Indirect(reflect.ValueOf(dst))
		value := reflect.ValueOf(data).Elem()
		dist.Set(value)
	default:
		return decodeMessage(dst.ProtoReflect(), data)
	}
	return nil
}

func decodeMessage(rv protoreflect.Message, data *pb.Record) error {
	md := rv.Descriptor()
	fds := md.Fields()
	if string(md.FullName()) != data.TypeName {
		return fmt.Errorf("%s: type mismatch: want %s, but got %s", pkgName, md.FullName(), data.TypeName)
	}
	for n, field := range data.Fields {
		fd := md.Fields().ByNumber(protoreflect.FieldNumber(n))
		if fd == nil {
			// TODO ログ出し
			continue
		}
		if err := decodeField(rv, fd, field.Value); err != nil {
			// TODO ログ出し
			continue
		}
	}
	pkName := codecinternal.GetPrimaryKeyName(string(md.FullName()))
	if fd := fds.ByName(protoreflect.Name(pkName)); fd != nil {
		rv.Set(fd, protoreflect.ValueOfString(data.RecordId))
	}
	if fd := fds.ByName(protoreflect.Name("created_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(data.CreatedAt))
	}
	if fd := fds.ByName(protoreflect.Name("updated_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(data.UpdatedAt))
	}
	if fd := fds.ByName(protoreflect.Name("deleted_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(data.DeletedAt))
	}
	if fd := fds.ByName(protoreflect.Name("shared_properties")); fd != nil && data.SharedProperties != nil {
		rv.Set(fd, protoreflect.ValueOfMessage(data.SharedProperties.ProtoReflect()))
	}
	return nil
}

func decodeField(rv protoreflect.Message, fd protoreflect.FieldDescriptor, data []byte) error {
	var value protoreflect.Value
	if fd.Cardinality() == protoreflect.Repeated {
		if fd.IsMap() {
			dst := rv.Mutable(fd).Map()
			dst.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
				dst.Clear(mk)
				return true
			})
			var m interface{}
			jd := json.NewDecoder(bytes.NewReader(data))
			jd.UseNumber()
			if err := jd.Decode(&m); err != nil {
				return err
			}
			rv := reflect.ValueOf(m)
			iter := rv.MapRange()
			for iter.Next() {
				k := iter.Key()
				v := iter.Value()
				key, err := decodeMapKey(k)
				if err != nil {
					return err
				}
				b, err := json.Marshal(v.Interface())
				if err != nil {
					return err
				}
				value, err := decodeValue(fd, b)
				if err != nil {
					return err
				}
				dst.Set(key, value)
			}
			value = protoreflect.ValueOf(dst)
		} else {
			var l []json.RawMessage
			if err := json.Unmarshal(data, &l); err != nil {
				return err
			}
			dst := rv.Mutable(fd).List()
			dst.Truncate(0)
			for i := range l {
				v, err := decodeValue(fd, l[i])
				if err != nil {
					return err
				}
				dst.Append(v)
			}
			value = protoreflect.ValueOf(dst)
		}
	} else {
		if v, err := decodeValue(fd, data); err != nil {
			return err
		} else {
			value = v
		}
	}
	if !value.IsValid() {
		return nil
	}
	rv.Set(fd, value)
	return nil
}

func decodeValue(fd protoreflect.FieldDescriptor, data []byte) (protoreflect.Value, error) {
	// refer to proto definition
	kind := fd.Kind()
	if fd.IsMap() {
		kind = fd.MapValue().Kind()
	}
	var val interface{}
	if kind == protoreflect.MessageKind {
		val = &pb.Record{}
	}
	jd := json.NewDecoder(bytes.NewReader(data))
	jd.UseNumber()
	if err := jd.Decode(&val); err != nil {
		return protoreflect.Value{}, err
	}
	switch kind {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(val.(bool)), nil
	case protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sfixed32Kind:
		n, err := strconv.ParseInt(string(val.(json.Number)), 10, 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfInt32(int32(n)), nil
	case protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind:
		n, err := strconv.ParseInt(string(val.(json.Number)), 10, 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfInt64(int64(n)), nil
	case protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind:
		n, err := strconv.ParseUint(string(val.(json.Number)), 10, 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfUint32(uint32(n)), nil
	case protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind:
		n, err := strconv.ParseUint(string(val.(json.Number)), 10, 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfUint64(uint64(n)), nil
	case protoreflect.FloatKind:
		n, err := strconv.ParseFloat(string(val.(json.Number)), 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfFloat32(float32(n)), nil
	case protoreflect.DoubleKind:
		n, err := strconv.ParseFloat(string(val.(json.Number)), 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfFloat64(float64(n)), nil
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(val.(string)), nil
	case protoreflect.BytesKind:
		// goのjsonパッケージは、[]byteをbase64した文字列に変換する
		decoded, err := base64.StdEncoding.DecodeString(val.(string))
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfBytes(decoded), nil
	case protoreflect.EnumKind:
		n, err := strconv.ParseInt(string(val.(json.Number)), 10, 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(n)), nil
	case protoreflect.MessageKind:
		if val == nil {
			return protoreflect.Value{}, nil
		}
		rec := val.(*pb.Record)
		// protoファイルの存在する状況で実行されるため、型名で探せば対応するメッセージが見つかるはず
		msgTyp, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(rec.TypeName))
		if err != nil {
			return protoreflect.Value{}, err
		}
		dst := msgTyp.New()
		if err := decodeMessage(dst, rec); err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfMessage(dst), nil
	// case protoreflect.GroupKind:
	default:
		return protoreflect.ValueOf(val), nil
	}
}

func decodeMapKey(key reflect.Value) (protoreflect.MapKey, error) {
	var val interface{}
	switch key.Kind() {
	case reflect.String:
		var mKey mapKey
		jd := json.NewDecoder(bytes.NewReader([]byte(key.String())))
		jd.UseNumber()
		if err := jd.Decode(&mKey); err != nil {
			return protoreflect.MapKey{}, err
		}
		switch mKey.Type {
		case reflect.String.String():
			val = (mKey.Value).(string)
		case reflect.Bool.String():
			val = (mKey.Value).(bool)
		case reflect.Int32.String():
			n, err := strconv.ParseInt(string(mKey.Value.(json.Number)), 10, 32)
			if err != nil {
				return protoreflect.MapKey{}, err
			}
			val = int32(n)
		case reflect.Int64.String():
			n, err := strconv.ParseInt(string(mKey.Value.(json.Number)), 10, 64)
			if err != nil {
				return protoreflect.MapKey{}, err
			}
			val = int64(n)
		case reflect.Uint32.String():
			n, err := strconv.ParseUint(string(mKey.Value.(json.Number)), 10, 32)
			if err != nil {
				return protoreflect.MapKey{}, err
			}
			val = uint32(n)
		case reflect.Uint64.String():
			n, err := strconv.ParseUint(string(mKey.Value.(json.Number)), 10, 64)
			if err != nil {
				return protoreflect.MapKey{}, err
			}
			val = uint64(n)
		default:
			val = key.String()
		}
	default:
		panic("decodeMapKey: map key must be string type")
	}
	v := protoreflect.ValueOf(val)
	return v.MapKey(), nil
}
