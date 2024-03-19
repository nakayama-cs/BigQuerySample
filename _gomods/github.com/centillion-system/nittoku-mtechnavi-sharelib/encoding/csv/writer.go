package csv

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
)

type Marshaler interface {
	MarshalCSV() ([]string, error)
}

type Writer struct {
	Comma rune

	UseCRLF bool

	ForceQuote bool

	w io.Writer

	err error

	once sync.Once
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Comma:   ',',
		UseCRLF: true,
		w:       w,
	}
}

func (w *Writer) Write(v Marshaler) error {
	record, err := v.MarshalCSV()
	if err != nil {
		panic(err)
	}
	return w.write(record)
}

// v is must be a slice of proto.Message
func (w *Writer) WriteAll(v any) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("v is must be a pointer to a slice: %T", v))
	}
	// elTyp expects a pointer to a struct assignable to Marshaler
	elTyp := rv.Type().Elem()
	if !elTyp.Implements(reflect.TypeOf((*Marshaler)(nil)).Elem()) {
		panic(fmt.Sprintf("v's elem must be a Marshaler: %v", elTyp))
	}
	for i := 0; i < rv.Len(); i++ {
		v := rv.Index(i).Interface().(Marshaler)
		if err := w.Write(v); err != nil {
			return err
		}
	}
	return nil
}

type StringMarshaler []string

func (l StringMarshaler) MarshalCSV() ([]string, error) {
	return l, nil
}

func (w *Writer) write(record []string) error {
	for i := range record {
		record[i] = w.quote(record[i])
	}
	if _, err := io.WriteString(w.w, strings.Join(record, string(w.Comma))); err != nil {
		return err
	}
	var err error
	if w.UseCRLF {
		_, err = io.WriteString(w.w, "\r\n")
	} else {
		_, err = io.WriteString(w.w, "\n")
	}
	return err
}

func (w *Writer) quote(s string) string {
	var b strings.Builder

	needQuote := w.ForceQuote
	for _, ch := range s {
		switch ch {
		case '"', '\r', '\n', w.Comma:
			needQuote = true
		}
		var err error
		if ch == '"' {
			_, err = b.WriteString(`""`)
		} else {
			_, err = b.WriteRune(ch)
		}
		if err != nil {
			panic(err)
		}
	}
	if needQuote {
		return `"` + b.String() + `"`
	} else {
		return b.String()
	}
}
