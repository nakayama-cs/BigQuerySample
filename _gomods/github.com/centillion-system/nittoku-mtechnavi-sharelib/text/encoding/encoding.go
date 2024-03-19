package encoding

import (
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var (
	UTF8 encoding.Encoding = &utf8Encoding{}

	ShiftJIS = &shiftjisEncoding{}
)

func NewReader(r io.Reader, enc encoding.Encoding) io.Reader {
	return transform.NewReader(r, enc.NewDecoder())
}

func NewWriter(w io.Writer, enc encoding.Encoding) io.Writer {
	return transform.NewWriter(w, enc.NewEncoder())
}
