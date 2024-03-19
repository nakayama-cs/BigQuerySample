package configbase

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvTag(t *testing.T) {
	cases := []struct {
		Tag     reflect.StructTag
		Name    string
		Options map[string][]interface{}
	}{
		{reflect.StructTag(``), "", map[string][]interface{}{
			"base": []interface{}{"", false},
		}},
		{reflect.StructTag(`env:""`), "", map[string][]interface{}{
			"base": []interface{}{"", false},
		}},
		{reflect.StructTag(`env:"XXX_YYY"`), "XXX_YYY", map[string][]interface{}{
			"base": []interface{}{"", false},
		}},
		{reflect.StructTag(`env:"XXX_YYY,base=60"`), "XXX_YYY", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", false},
		}},
		{reflect.StructTag(`env:",base=60"`), "", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", false},
		}},
		{reflect.StructTag(`env:"XXX_YYY,base=60,switch"`), "XXX_YYY", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", true},
		}},
		{reflect.StructTag(`env:",base=60,switch"`), "", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", true},
		}},
		{reflect.StructTag(`env:"XXX_YYY,switch,base=60"`), "XXX_YYY", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", true},
		}},
		{reflect.StructTag(`env:",switch,base=60"`), "", map[string][]interface{}{
			"base":   []interface{}{"60", true},
			"xxx":    []interface{}{"", false},
			"switch": []interface{}{"", true},
		}},
	}
	for _, c := range cases {
		tag := envTagFromStructTag(c.Tag)
		assert.Equalf(t, c.Name, tag.Name(), "%q", c.Tag)
		for optName, opt := range c.Options {
			v, ok := tag.Option(optName)
			assert.Equalf(t, opt[0], v, "%q", c.Tag)
			assert.Equalf(t, opt[1], ok, "%q", c.Tag)
		}
	}
}
