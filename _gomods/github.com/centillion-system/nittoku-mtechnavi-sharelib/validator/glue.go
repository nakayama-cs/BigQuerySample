package validator

import (
	"mtechnavi/sharelib/constants"
	"reflect"
	"regexp"
	"strings"
)

// このファイルには、生成されたコードを元にして処理する内容をまとめている。
// 生成コードに変更がある度にメンテが必要となる。

type validationError interface {
	Field() string

	Reason() string

	Cause() error

	Key() bool

	ErrorName() string
}

type translator struct {
	Req interface{}
}

func (t *translator) Field(field string) string {
	// protoc-gen-validateはGoのstruct field名を返してくるため、protocが付与するstruct tagから元の名前を解決する
	rt := reflect.TypeOf(t.Req)
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	rf, ok := rt.FieldByName(field)
	if !ok {
		return field
	}
	s, ok := rf.Tag.Lookup("protobuf")
	if !ok {
		return field
	}
	for _, item := range strings.Split(s, ",") {
		if !strings.HasPrefix(item, "name=") {
			continue
		}
		return strings.TrimPrefix(item, "name=")
	}
	return field
}

func (t *translator) Reason(reason string) (string, []string) {
	patternE0000017 := regexp.MustCompile("value length must be between ([0-9]+) and ([0-9]+) runes, inclusive")

	switch {
	case patternE0000017.Match([]byte(reason)):
		gs := patternE0000017.FindSubmatch([]byte(reason))
		return constants.MessageName_E0000017.String(), []string{string(gs[1]), string(gs[2])}
	}

	return "", nil
}
