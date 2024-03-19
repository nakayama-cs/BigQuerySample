package dataproxy

import (
	recordpb "mtechnavi/sharelib/protobuf/record"
)

func Example() {
	var l []*recordpb.Record

	l2 := ToMessages(l...)
	l = FromMessages[*recordpb.Record](l2...)
}
