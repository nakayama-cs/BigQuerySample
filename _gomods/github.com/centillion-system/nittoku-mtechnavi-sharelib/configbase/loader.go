package configbase

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/stoewer/go-strcase"
)

// Load は指定された取得元から、 v で渡された struct の各フィールドに対して値を取得し、設定する。
func Load(ctx context.Context, v interface{}) error {
	var rv reflect.Value
	if val, ok := v.(reflect.Value); ok {
		rv = val
	} else {
		rv = reflect.ValueOf(v)
	}
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("v is not a struct")
	}
	for i := 0; i < rv.NumField(); i++ {
		rf := rv.Field(i)
		tf := rv.Type().Field(i)
		if !tf.IsExported() {
			continue
		}
		if err := loadField(ctx, rf, tf); err != nil {
			return err
		}
	}
	return nil
}

func loadField(ctx context.Context, rf reflect.Value, tf reflect.StructField) error {
	if rf.Kind() == reflect.Pointer && rf.IsNil() {
		rf.Set(reflect.New(rf.Type().Elem()))
	}
	for rf.Kind() == reflect.Pointer {
		rf = rf.Elem()
	}
	tag := envTagFromStructTag(tf.Tag)
	envName := tag.Name()
	if envName == "-" {
		return nil
	} else if envName == "" {
		envName = strcase.UpperSnakeCase(tf.Name)
	}
	prefix := pathFromContext(ctx)
	if prefix != "" {
		prefix += "_"
	}
	envName = prefix + envName

	envValue, err := getEnvValue(envName)
	if err != nil {
		return fmt.Errorf("failed to get environment variable %q: %w", envName, err)
	}
	if envValue == "" {
		envValue, _ = tag.Option("default")
	}
	switch rf.Kind() {
	case reflect.String:
		rf.SetString(envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := toInt(envValue, tf)
		if err != nil {
			return err
		}
		rf.SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := toUint(envValue, tf)
		if err != nil {
			return err
		}
		rf.SetUint(v)
	case reflect.Bool:
		v, err := toBool(envValue, tf)
		if err != nil {
			return err
		}
		rf.SetBool(v)
	case reflect.Struct:
		ctx = contextWithPath(ctx, envName)
		if err := Load(ctx, rf); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unable to convert string to %s of field %s", rf.Kind(), tf.Name)
	}
	return nil
}

func getEnvValue(envName string) (string, error) {
	// ファイルからの取得指定
	if s := os.Getenv(envName + "_FROM_FILE"); s != "" {
		data, err := os.ReadFile(s)
		if err != nil {
			return "", fmt.Errorf("failed to read from file: %v: %w", s, err)
		}
		return strings.TrimRight(string(data), "\r\n"), nil
	}
	return os.Getenv(envName), nil
}
