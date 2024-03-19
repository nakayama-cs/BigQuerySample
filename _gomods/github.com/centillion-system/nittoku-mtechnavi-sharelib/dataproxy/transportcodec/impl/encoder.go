package impl

import (
	"encoding/json"
	"fmt"
	"mtechnavi/sharelib/dataproxy/codec"
	sharelibpb "mtechnavi/sharelib/protobuf"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	recordpb "mtechnavi/sharelib/protobuf/record"
	"path"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Encode(src proto.Message, resolver protodesc.Resolver) (*recordpb.Record, error) {
	if resolver == nil {
		resolver = protoregistry.GlobalFiles
	}

	rv := src.ProtoReflect()
	fs, err := newFileDescriptorSet(src.ProtoReflect().Descriptor(), resolver)
	if err != nil {
		return nil, err
	}

	var out recordpb.Record
	out.TypeName = string(rv.Descriptor().FullName())
	if v, ok := getFieldValueByName(rv, codec.GetPrimaryKeyName(out.TypeName)); ok {
		out.RecordId = v.String()
	}
	if v, ok := getFieldValueByName(rv, "shared_properties"); ok {
		out.SharedProperties = toEmbeddedSharedProperties(v.Message())
	}
	if v, ok := getFieldValueByName(rv, "created_at"); ok {
		out.CreatedAt = v.Int()
	}
	if v, ok := getFieldValueByName(rv, "updated_at"); ok {
		out.UpdatedAt = v.Int()
	}
	if v, ok := getFieldValueByName(rv, "deleted_at"); ok {
		out.DeletedAt = v.Int()
	}
	out.Fields = make(map[int32]*recordpb.Field, 3)
	out.Fields[fieldNumberFormat] = &recordpb.Field{
		Value: []byte(formatVersion),
	}
	if b, err := proto.Marshal(fs); err != nil {
		return nil, err
	} else {
		out.Fields[fieldNumberDescriptor] = &recordpb.Field{
			Value: b,
		}
	}
	if b, err := proto.Marshal(src); err != nil {
		return nil, err
	} else {
		out.Fields[fieldNumberPayload] = &recordpb.Field{
			Value: b,
		}
	}
	if l, err := createIndexs(rv); err != nil {
		return nil, err
	} else {
		out.Indexs = l
	}
	return &out, nil
}

func toEmbeddedSharedProperties(rv protoreflect.Message) *sharelibpb.EmbeddedSharedProperties {
	if v, ok := rv.Interface().(*sharelibpb.EmbeddedSharedProperties); ok {
		return v
	}

	// Decodeしたものを再度Encodeにかけたときなど、 *dynamicpb.Message となるため、変換必須
	var out sharelibpb.EmbeddedSharedProperties
	rv.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		outRv := out.ProtoReflect()
		outFd := outRv.Descriptor().Fields().ByTextName(fd.TextName())
		if outFd == nil {
			return true
		}
		outRv.Set(outFd, v)
		return true
	})
	return &out
}

func newFileDescriptorSet(md protoreflect.MessageDescriptor, resolver protodesc.Resolver) (*descriptorpb.FileDescriptorSet, error) {
	fileDescs, err := gatherFileDescriptors(protodesc.ToFileDescriptorProto(md.ParentFile()), resolver)
	if err != nil {
		return nil, err
	}

	var out descriptorpb.FileDescriptorSet
	out.File = make([]*descriptorpb.FileDescriptorProto, 0, len(fileDescs))
	for _, v := range fileDescs {
		out.File = append(out.File, v)
	}
	return &out, nil
}

func gatherFileDescriptors(v *descriptorpb.FileDescriptorProto, resolver protodesc.Resolver) (map[string]*descriptorpb.FileDescriptorProto, error) {
	out := map[string]*descriptorpb.FileDescriptorProto{}
	// "{{package}}/{{filename}}" でユニーク化し、多重収集を防止
	out[path.Join(v.GetPackage(), v.GetName())] = v
	for _, dep := range v.Dependency {
		depFile, err := resolver.FindFileByPath(dep)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to find file: %s: %w", pkgName, dep, err)
		}
		val, err := gatherFileDescriptors(protodesc.ToFileDescriptorProto(depFile), resolver)
		if err != nil {
			return nil, err
		}
		for k := range val {
			out[k] = val[k]
		}
	}
	return out, nil
}

func getFieldValueByName(rv protoreflect.Message, name string) (protoreflect.Value, bool) {
	fd := rv.Descriptor().Fields().ByName(protoreflect.Name(name))
	if fd == nil {
		return protoreflect.Value{}, false
	}
	return rv.Get(fd), true
}

func createIndexs(rv protoreflect.Message) ([]*recordpb.Index, error) {
	var out []*recordpb.Index
	fds := rv.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
		// ネストしたメッセージフィールドのIndexは不可
		fd := fds.Get(i)
		_, do := getFieldExtension(fd)
		if do == nil {
			continue
		}
		// XXX: プリミティブのみがIndex登録されることが想定
		b, err := json.Marshal(rv.Get(fd).Interface())
		if err != nil {
			return nil, err
		}
		out = append(out, &recordpb.Index{
			Name:  do.Index.Name,
			Value: b,
		})
	}
	return out, nil
}

func getFieldExtension(fd protoreflect.FieldDescriptor) (*mtnpb.VisibilityOptions, *mtnpb.DataproxyOptions) {
	var (
		vo *mtnpb.VisibilityOptions
		do *mtnpb.DataproxyOptions
	)
	proto.RangeExtensions(fd.Options(), func(xt protoreflect.ExtensionType, ev interface{}) bool {
		if v, ok := ev.(*mtnpb.VisibilityOptions); ok {
			vo = v
		}
		if v, ok := ev.(*mtnpb.DataproxyOptions); ok {
			do = v
		}
		return true
	})
	return vo, do
}
