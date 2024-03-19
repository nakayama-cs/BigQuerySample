package configbase

import (
	"reflect"
	"strconv"
	"strings"
)

func toInt(s string, tf reflect.StructField) (int64, error) {
	if s == "" {
		return 0, nil
	}
	tag := envTagFromStructTag(tf.Tag)
	baseStr, _ := tag.Option("base")
	if baseStr == "" {
		baseStr = "10"
	}
	base, err := strconv.ParseInt(baseStr, 10, 32)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseInt(s, int(base), 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func toUint(s string, tf reflect.StructField) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	tag := envTagFromStructTag(tf.Tag)
	baseStr, _ := tag.Option("base")
	if baseStr == "" {
		baseStr = "10"
	}
	base, err := strconv.ParseInt(baseStr, 10, 32)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseUint(s, int(base), 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func toBool(s string, tf reflect.StructField) (bool, error) {
	return strings.ToLower(s) == "true", nil
}
