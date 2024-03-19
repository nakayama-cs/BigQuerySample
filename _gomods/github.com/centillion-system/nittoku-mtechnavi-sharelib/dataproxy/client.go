// dataproxy パッケージは、dataproxyを利用しやすくするためのラッパークライアントを提供するパッケージ。
package dataproxy

import (
	"fmt"
	dataproxypb "mtechnavi/dataproxy/protobuf"
	"mtechnavi/sharelib/dataproxy/transportcodec"
	recordpb "mtechnavi/sharelib/protobuf/record"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Client struct {
	c dataproxypb.DataproxyClient
}

func New(cc grpc.ClientConnInterface) *Client {
	return &Client{
		c: dataproxypb.NewDataproxyClient(cc),
	}
}

func Encode(l []proto.Message) ([]*recordpb.Record, error) {
	out := make([]*recordpb.Record, len(l))
	for i := range l {
		if v, err := transportcodec.Encode(l[i], nil); err != nil {
			return nil, err
		} else {
			out[i] = v
		}
	}
	return out, nil
}

func Decode(dst []proto.Message, src []*recordpb.Record) error {
	if len(dst) != len(src) {
		panic(fmt.Sprintf("length mismatched: len(dst) == %d, len(src) == %d", len(dst), len(src)))
	}
	for i := range src {
		v, err := transportcodec.Decode(src[i])
		if err != nil {
			return err
		}
		rv := dst[i].ProtoReflect()
		rv.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
			rv.Clear(fd)
			return true
		})
		v.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, fv protoreflect.Value) bool {
			rv.Set(fd, fv)
			return true
		})
	}
	return nil
}
