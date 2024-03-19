package transportcodec

import (
	"errors"
	"fmt"
	legacycodec "mtechnavi/sharelib/dataproxy/codec"
	"mtechnavi/sharelib/dataproxy/transportcodec/impl"
	recordpb "mtechnavi/sharelib/protobuf/record"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const (
	pkgName = "mtechnavi/dataproxy/transportcodec"
)

type EncodeConfig struct {
	Resolver protodesc.Resolver

	ResolverFromRecord *recordpb.Record

	FormatFromRecord *recordpb.Record
}

func (ec *EncodeConfig) getResolver() (protodesc.Resolver, error) {
	if ec == nil {
		return nil, nil
	}
	if ec.Resolver != nil && ec.ResolverFromRecord != nil {
		panic(fmt.Sprintf("Resolver or ResolverFromRecord are exclusive"))
	}
	if ec.Resolver != nil {
		return ec.Resolver, nil
	}
	// DataProxy側の処理を定型化したいので、v0フォーマットのときは無視する
	if ec.ResolverFromRecord != nil && impl.IsCompat(ec.ResolverFromRecord) {
		r, err := impl.NewResolverFromRecord(ec.ResolverFromRecord)
		if err != nil {
			return nil, fmt.Errorf("%s: invalid ResolverFromRecord: %w", pkgName, err)
		}
		return r, nil
	}
	return nil, nil
}

func (ec *EncodeConfig) useLegacyCodec() bool {
	if ec == nil || ec.FormatFromRecord == nil {
		return false
	}
	return !impl.IsCompat(ec.FormatFromRecord)
}

func Encode(src proto.Message, cfg *EncodeConfig) (*recordpb.Record, error) {
	if cfg.useLegacyCodec() {
		return legacycodec.Encode(src)
	}

	var (
		resolver protodesc.Resolver
	)
	if v, err := cfg.getResolver(); err != nil {
		return nil, err
	} else {
		resolver = v
	}
	return impl.Encode(src, resolver)
}

func IsLegacy(v *recordpb.Record) bool {
	return !impl.IsCompat(v)
}

func Decode(src *recordpb.Record) (proto.Message, error) {
	if impl.IsCompat(src) {
		return impl.Decode(src)
	}

	// legacy support
	// クライアント側でのデコード時には、翻訳単位内に型情報があるはず
	var (
		typName = protoreflect.FullName(src.TypeName)
		m       proto.Message
	)
	if typ, err := protoregistry.GlobalTypes.FindMessageByName(typName); errors.Is(err, protoregistry.NotFound) {
		// 型情報がないなら、DataProxy上でのデコード
		m = &recordpb.Record{}
	} else if err != nil {
		return nil, err
	} else {
		m = typ.New().Interface()
	}
	if err := legacycodec.Decode(m, src); err != nil {
		return nil, err
	}
	return m, nil
}
