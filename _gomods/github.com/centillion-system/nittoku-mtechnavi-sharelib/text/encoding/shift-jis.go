package encoding

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type shiftjisEncoding struct{}

func (*shiftjisEncoding) NewEncoder() *encoding.Encoder {
	e := japanese.ShiftJIS.NewEncoder()
	// ShiftJISで表現できない文字を正規化して出力する
	e.Transformer = transform.Chain(norm.NFC, e.Transformer)
	return e
}

func (*shiftjisEncoding) NewDecoder() *encoding.Decoder {
	d := japanese.ShiftJIS.NewDecoder()
	return d
}
