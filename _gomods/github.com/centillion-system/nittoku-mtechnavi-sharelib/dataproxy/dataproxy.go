package dataproxy

import (
	"google.golang.org/protobuf/proto"
)

func ToMessages[M proto.Message](l ...M) []proto.Message {
	out := make([]proto.Message, len(l))
	for i := range l {
		out[i] = l[i]
	}
	return out
}

func FromMessages[M proto.Message](l ...proto.Message) []M {
	out := make([]M, len(l))
	for i := range l {
		out[i] = l[i].(M)
	}
	return out
}
