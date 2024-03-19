package validator

import (
	"encoding/json"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/validator/protobuf"
	testdatapb "mtechnavi/sharelib/validator/testdata/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	var req testdatapb.Strings
	req.MaxLen_5 = "123456"
	req.MaxBytes_5 = "123456"
	req.NotContains = "foo"
	err := validate(&req)
	if !assert.Error(t, err) {
		return
	}
	n := req.ProtoReflect().Descriptor().Fields().Len()
	merr := err.(testdatapb.StringsMultiError)
	assert.Equal(t, n, len(merr), "num errors")
}

func TestToValidationError(t *testing.T) {
	var req testdatapb.Strings
	req.MaxLen_5 = "123456"
	req.MaxBytes_5 = "123456"
	req.NotContains = "foo"
	err := validate(&req)
	got := toValidationError(&req, err)
	want := &pb.ValidationError{
		Descriptors: []*pb.Descriptor{
			&pb.Descriptor{
				Field: "const_foo",
				// ErrorMessage: `invalid Strings.ConstFoo: value must equal foo`,
			},
			&pb.Descriptor{
				Field: "len_5",
				// ErrorMessage: `invalid Strings.Len_5: value length must be 5 runes`,
			},
			&pb.Descriptor{
				Field: "min_len_5",
				// ErrorMessage: `invalid Strings.MinLen_5: value length must be at least 5 runes`,
			},
			&pb.Descriptor{
				Field: "max_len_5",
				// ErrorMessage: `invalid Strings.MaxLen_5: value length must be at most 5 runes`,
			},
			&pb.Descriptor{
				Field:       "between_1_5",
				MessageName: constants.MessageName_E0000017.String(),
				MessageArgs: []string{"1", "5"},
				// ErrorMessage: `invalid Strings.Between_1_5: value length must be between 1 and 5 runes, inclusive`,
			},
			&pb.Descriptor{
				Field: "min_bytes_5",
				// ErrorMessage: `invalid Strings.MinBytes_5: value length must be at least 5 bytes`,
			},
			&pb.Descriptor{
				Field: "max_bytes_5",
				// ErrorMessage: `invalid Strings.MaxBytes_5: value length must be at most 5 bytes`,
			},
			&pb.Descriptor{
				Field: "between_bytes_1_5",
				// ErrorMessage: `invalid Strings.BetweenBytes_1_5: value length must be between 1 and 5 bytes, inclusive`,
			},
			&pb.Descriptor{
				Field: "pattern",
				// ErrorMessage: `invalid Strings.Pattern: value does not match regex pattern "re2-compliant"`,
			},
			&pb.Descriptor{
				Field: "prefix",
				// ErrorMessage: `invalid Strings.Prefix: value does not have prefix "foo"`,
			},
			&pb.Descriptor{
				Field: "suffix",
				// ErrorMessage: `invalid Strings.Suffix: value does not have suffix "foo"`,
			},
			&pb.Descriptor{
				Field: "contains",
				// ErrorMessage: `invalid Strings.Contains: value does not contain substring "foo"`,
			},
			&pb.Descriptor{
				Field: "not_contains",
				// ErrorMessage: `invalid Strings.NotContains: value contains substring "foo"`,
			},
		},
	}
	assert.JSONEq(t, jsonStr(want), jsonStr(got))
}

func jsonPP(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func jsonStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
