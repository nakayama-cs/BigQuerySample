package impl

import (
	"fmt"
	recordpb "mtechnavi/sharelib/protobuf/record"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func Decode(src *recordpb.Record) (proto.Message, error) {
	resolver, err := NewResolverFromRecord(src)
	if err != nil {
		return nil, err
	}

	var dataBytes []byte
	if v, ok := src.Fields[fieldNumberPayload]; ok {
		dataBytes = v.Value
	} else {
		return nil, fmt.Errorf("%s: invalid format: message binary at field %d is not found.", pkgName, fieldNumberPayload)
	}

	var md protoreflect.MessageDescriptor
	if v, err := resolver.FindDescriptorByName(protoreflect.FullName(src.TypeName)); err != nil {
		return nil, fmt.Errorf("%s: message %s is not containing given runtime descriptor: %w", pkgName, src.TypeName, err)
	} else {
		md = v.(protoreflect.MessageDescriptor)
	}
	dst := dynamicpb.NewMessageType(md).New().Interface()

	if err := proto.Unmarshal(dataBytes, dst); err != nil {
		return nil, err
	}
	return dst, nil
}

func NewResolverFromRecord(src *recordpb.Record) (protodesc.Resolver, error) {
	var descBytes []byte
	if v, ok := src.Fields[fieldNumberDescriptor]; ok {
		descBytes = v.Value
	} else {
		return nil, fmt.Errorf("%s: invalid format: runtime descriptor at field %d is not found.", pkgName, fieldNumberDescriptor)
	}

	var fs descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(descBytes, &fs); err != nil {
		return nil, fmt.Errorf("%s: invalid runtime descriptor: %w", pkgName, err)
	}
	files, err := protodesc.NewFiles(&fs)
	if err != nil {
		return nil, err
	}
	return files, nil
}
