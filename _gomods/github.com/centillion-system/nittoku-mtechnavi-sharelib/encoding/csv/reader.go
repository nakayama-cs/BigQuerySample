package csv

import (
	gocsv "encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/hashicorp/go-multierror"
)

type Unmarshaler interface {
	UnmarshalCSV([]string) error
}

type Reader struct {
	*gocsv.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		Reader: gocsv.NewReader(r),
	}
}

func (r *Reader) Read(v Unmarshaler) error {
	data, err := r.Reader.Read()
	if err != nil {
		return err
	}
	if err = v.UnmarshalCSV(data); err != nil {
		switch err := err.(type) {
		case *Error:
			for _, cerr := range err.Causes {
				switch e := cerr.(type) {
				case *Error:
					if e.Column > 0 {
						e.Line, _ = r.FieldPos(e.Column - 1)
					}
				}
			}
			return err
		case *multierror.Error:
			for _, cerr := range err.WrappedErrors() {
				switch e := cerr.(type) {
				case *Error:
					if e.Column > 0 {
						e.Line, _ = r.FieldPos(e.Column - 1)
					}
				}
			}
			return err
		default:
			return err
		}
	}
	return nil
}

// v is must be a pointer to a slice of Unmarshaler
func (r *Reader) ReadAll(v any) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("v is must be a pointer to a slice: %T", v))
	}
	// elTyp expects a pointer to a struct assignable to Unmarshaler
	elTyp := rv.Type().Elem()
	if !elTyp.Implements(reflect.TypeOf((*Unmarshaler)(nil)).Elem()) {
		panic(fmt.Sprintf("v's elem must be a Unmarshaler: %v", elTyp))
	}
	for {
		val := reflect.New(elTyp.Elem()).Interface().(Unmarshaler)

		if err := r.Read(val); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}
		rv = reflect.Append(rv, reflect.ValueOf(val))
	}
	reflect.ValueOf(v).Elem().Set(rv)
	return nil
}

type HeaderUnmarshaler struct{}

func (s *HeaderUnmarshaler) UnmarshalCSV([]string) error {
	return nil
}
