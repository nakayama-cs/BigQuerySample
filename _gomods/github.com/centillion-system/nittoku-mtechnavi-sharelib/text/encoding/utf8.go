package encoding

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// goは標準でutf8のため、バイパスのみ
type utf8Encoding struct{}

func (*utf8Encoding) NewEncoder() *encoding.Encoder {
	var e encoding.Encoder
	e.Transformer = transform.Nop
	return &e
}

func (*utf8Encoding) NewDecoder() *encoding.Decoder {
	var d encoding.Decoder
	d.Transformer = transform.Nop
	return &d
}
