package internal

import (
	"strings"
	"unicode"
)

func GetPrimaryKeyName(typename string) string {
	if typename == "" {
		return ""
	}
	typeslice := strings.Split(typename, ".")
	t := typeslice[len(typeslice)-1]
	idname := strings.ToLower(t[:1]) + t[1:] + "_id"
	return toSnakeCase(idname)
}

func toSnakeCase(s string) string {
	b := &strings.Builder{}
	for i, r := range s {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
