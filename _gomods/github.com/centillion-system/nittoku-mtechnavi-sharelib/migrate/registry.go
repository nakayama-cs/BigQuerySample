package migrate

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	registryMu sync.Mutex

	registry = map[protoreflect.FullName]proto.Message{}
)

func RegisterMessage(msg proto.Message) {
	rv := reflect.ValueOf(msg)
	if rv.IsNil() {
		typ := rv.Type()
		for typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
		}
		msg = reflect.New(typ).Interface().(proto.Message)
	}
	fullName := proto.MessageName(msg)

	registryMu.Lock()
	defer registryMu.Unlock()

	if _, dup := registry[fullName]; dup {
		panic("duplicated: " + fullName)
	}
	registry[fullName] = msg
}

func newMessage(fullName protoreflect.FullName) proto.Message {
	msg, ok := registry[fullName]
	if !ok {
		panic("not found: " + fullName)
	}
	return proto.Clone(msg)
}
