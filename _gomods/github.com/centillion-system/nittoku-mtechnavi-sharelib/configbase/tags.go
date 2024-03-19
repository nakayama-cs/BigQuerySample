package configbase

import (
	"reflect"
	"strings"
)

// `env:"name"`"
// `env:"name,key=name,key,..."`"
type envTag []string

func envTagFromStructTag(tag reflect.StructTag) envTag {
	return envTag(strings.Split(tag.Get("env"), ","))
}

func (l envTag) Name() string {
	if len(l) == 0 {
		return ""
	}
	return l[0]
}

func (l envTag) Option(key string) (string, bool) {
	opts := l[1:]
	// find "key=value"
	for _, s := range opts {
		if !strings.HasPrefix(s, key+"=") {
			continue
		}
		value := strings.TrimPrefix(s, key+"=")
		return value, true
	}
	// find "key"
	for _, s := range opts {
		if s != key {
			continue
		}
		return "", true
	}
	return "", false
}
