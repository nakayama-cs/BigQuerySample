package transportcodec

import (
	legacycodec "mtechnavi/sharelib/dataproxy/codec"
	pb "mtechnavi/sharelib/dataproxy/codec/testdata"
	"mtechnavi/sharelib/dataproxy/transportcodec/impl"
	recordpb "mtechnavi/sharelib/protobuf/record"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestEncode(t *testing.T) {
	src := &pb.ScalarTypes{
		String_: "hi",
	}
	dst, err := Encode(src, nil)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, string(proto.MessageName(src)), dst.TypeName)
	assert.Equal(t, true, impl.IsCompat(dst))
}

func TestDecode(t *testing.T) {
	t.Run("v0 on client-side", func(t *testing.T) {
		src, err := legacycodec.Encode(&pb.ScalarTypes{
			String_: "hi",
		})
		if err != nil {
			panic(err)
		}
		dst, err := Decode(src)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, src.TypeName, string(proto.MessageName(dst)))
	})
	t.Run("v0 on dataproxy-side", func(t *testing.T) {
		src := &recordpb.Record{
			TypeName: "mtn.testing.NotExistMessage",
			Fields: map[int32]*recordpb.Field{
				1: &recordpb.Field{
					Name:  "text",
					Value: []byte(`"hi"`),
				},
			},
		}
		dst_, err := Decode(src)
		if !assert.NoError(t, err) {
			return
		}
		dst, ok := dst_.(*recordpb.Record)
		if !ok {
			assert.FailNow(t, "Decode does not return a record")
		}
		assert.Equal(t, src.TypeName, dst.TypeName)
	})
	t.Run("v1 on client-side", func(t *testing.T) {
		src, err := Encode(&pb.ScalarTypes{
			String_: "hi",
		}, nil)
		if err != nil {
			panic(err)
		}
		dst, err := Decode(src)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, src.TypeName, string(proto.MessageName(dst)))
	})
	t.Run("v1 on dataproxy-side", func(t *testing.T) {
		src, err := Encode(&pb.ScalarTypes{
			String_: "hi",
		}, nil)
		if err != nil {
			panic(err)
		}
		_globalFiles := protoregistry.GlobalFiles
		protoregistry.GlobalFiles = nil
		defer func() {
			protoregistry.GlobalFiles = _globalFiles
		}()

		dst, err := Decode(src)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, src.TypeName, string(proto.MessageName(dst)))
	})
}
