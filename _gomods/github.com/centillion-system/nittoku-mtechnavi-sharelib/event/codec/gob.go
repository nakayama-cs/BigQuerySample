package codec

import (
	"bytes"
	"encoding/gob"
)

type Gob struct{}

func (*Gob) Encode(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (*Gob) Decode(data []byte, v interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}
