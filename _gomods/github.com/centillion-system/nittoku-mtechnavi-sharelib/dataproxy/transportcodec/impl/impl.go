package impl

import (
	recordpb "mtechnavi/sharelib/protobuf/record"

	"golang.org/x/mod/semver"
)

const (
	pkgName = "dataproxy/transportcodec/impl"

	formatVersion = "v1.0.0"

	fieldNumberFormat     = 0
	fieldNumberDescriptor = 1
	fieldNumberPayload    = 2
)

func IsCompat(v *recordpb.Record) bool {
	f := v.Fields[fieldNumberFormat]
	if f == nil {
		return false
	}
	verStr := string(f.Value)
	return semver.Major(verStr) == "v1"
}
