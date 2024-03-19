package codec_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mtechnavi/sharelib/dataproxy/codec"
	testdatapb "mtechnavi/sharelib/dataproxy/codec/testdata"
	sharelibpb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/protobuf/mtn"
	recordpb "mtechnavi/sharelib/protobuf/record"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

const (
	epsilon = 0.000001
)

func TestScalarTypes(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		const typeName = "mtechnavi.testdata.ScalarTypes"
		var src testdatapb.ScalarTypes

		expects := map[int32]*recordpb.Field{
			1:  {Name: "double_", Value: jsonBytes(float64(0)), Visibility: &mtn.VisibilityOptions{}},
			3:  {Name: "float_", Value: jsonBytes(float32(0)), Visibility: &mtn.VisibilityOptions{}},
			5:  {Name: "int32_", Value: jsonBytes(int32(0)), Visibility: &mtn.VisibilityOptions{}},
			7:  {Name: "int64_", Value: jsonBytes(int64(0)), Visibility: &mtn.VisibilityOptions{}},
			9:  {Name: "uint32_", Value: jsonBytes(uint32(0)), Visibility: &mtn.VisibilityOptions{}},
			11: {Name: "uint64_", Value: jsonBytes(uint64(0)), Visibility: &mtn.VisibilityOptions{}},
			13: {Name: "sint32_", Value: jsonBytes(int32(0)), Visibility: &mtn.VisibilityOptions{}},
			15: {Name: "sint64_", Value: jsonBytes(int64(0)), Visibility: &mtn.VisibilityOptions{}},
			17: {Name: "fixed32_", Value: jsonBytes(uint32(0)), Visibility: &mtn.VisibilityOptions{}},
			19: {Name: "fixed64_", Value: jsonBytes(uint64(0)), Visibility: &mtn.VisibilityOptions{}},
			21: {Name: "sfixed32_", Value: jsonBytes(int32(0)), Visibility: &mtn.VisibilityOptions{}},
			23: {Name: "sfixed64_", Value: jsonBytes(int64(0)), Visibility: &mtn.VisibilityOptions{}},
			25: {Name: "bool_", Value: jsonBytes(bool(false)), Visibility: &mtn.VisibilityOptions{}},
			27: {Name: "string_", Value: jsonBytes(string("")), Visibility: &mtn.VisibilityOptions{}},
			29: {Name: "bytes_", Value: jsonBytes([]byte{}), Visibility: &mtn.VisibilityOptions{}},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, typeName, rec.TypeName)
		assert.Len(t, rec.Fields, len(expects))
		assert.Equal(t, expects, rec.Fields)
	})
	t.Run("single",
		func(t *testing.T) {
			const typeName = "mtechnavi.testdata.ScalarTypes"
			var src testdatapb.ScalarTypes
			src.Double_ = 1.1
			src.Float_ = 2.1
			src.Int32_ = 3
			src.Int64_ = 4
			src.Uint32_ = 5
			src.Uint64_ = 6
			src.Sint32_ = 7
			src.Sint64_ = 8
			src.Fixed32_ = 9
			src.Fixed64_ = 10
			src.Sfixed32_ = 11
			src.Sfixed64_ = 12
			src.Bool_ = true
			src.String_ = "14"
			src.Bytes_ = []byte{0xca, 0xfe, 0xba, 0xbe}

			expects := []struct {
				Number int
				Name   string
				JSON   string
			}{
				{1, "double_", `1.1`},
				{3, "float_", `2.1`},
				{5, "int32_", `3`},
				{7, "int64_", `4`},
				{9, "uint32_", `5`},
				{11, "uint64_", `6`},
				{13, "sint32_", `7`},
				{15, "sint64_", `8`},
				{17, "fixed32_", `9`},
				{19, "fixed64_", `10`},
				{21, "sfixed32_", `11`},
				{23, "sfixed64_", `12`},
				{25, "bool_", `true`},
				{27, "string_", `"14"`},
				{29, "bytes_", b64bytes(src.Bytes_)},
			}

			rec, err := codec.Encode(&src)
			if !assert.NoError(t, err) {
				return
			}

			assert.Equal(t, typeName, rec.TypeName)
			assert.Len(t, rec.Fields, 15)
			for _, exp := range expects {
				field := rec.Fields[int32(exp.Number)]
				assert.Equal(t, exp.Name, field.Name)
				switch exp.Number {
				case 1, 3:
					assert.InDeltaf(t,
						jsonNumberAsFloat64(exp.JSON), jsonNumberAsFloat64(string(field.Value)), epsilon,
						"field=%q", exp.Name)
				default:
					assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
				}
			}

			// codec.Debug(os.Stderr, rec)

			// decode from encoded
			var dst testdatapb.ScalarTypes
			if assert.NoError(t, codec.Decode(&dst, rec)) {
				assert.Exactly(t, &src, &dst)
			}
		})
	t.Run("repeated", func(t *testing.T) {
		const typeName = "mtechnavi.testdata.RepeatedTypes"
		var src testdatapb.RepeatedTypes
		src.RepeatedDouble_ = []float64{1.1}
		src.RepeatedFloat_ = []float32{2.1}
		src.RepeatedInt32_ = []int32{3}
		src.RepeatedInt64_ = []int64{4}
		src.RepeatedUint32_ = []uint32{5}
		src.RepeatedUint64_ = []uint64{6}
		src.RepeatedSint32_ = []int32{7}
		src.RepeatedSint64_ = []int64{8}
		src.RepeatedFixed32_ = []uint32{9}
		src.RepeatedFixed64_ = []uint64{10}
		src.RepeatedSfixed32_ = []int32{11}
		src.RepeatedSfixed64_ = []int64{12}
		src.RepeatedBool_ = []bool{true}
		src.RepeatedString_ = []string{"14"}
		src.RepeatedBytes_ = [][]byte{
			{0xca, 0xfe, 0xba, 0xbe},
		}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{101, "repeated_double_", `[1.1]`},
			{103, "repeated_float_", `[2.1]`},
			{105, "repeated_int32_", `[3]`},
			{107, "repeated_int64_", `[4]`},
			{109, "repeated_uint32_", `[5]`},
			{111, "repeated_uint64_", `[6]`},
			{113, "repeated_sint32_", `[7]`},
			{115, "repeated_sint64_", `[8]`},
			{117, "repeated_fixed32_", `[9]`},
			{119, "repeated_fixed64_", `[10]`},
			{121, "repeated_sfixed32_", `[11]`},
			{123, "repeated_sfixed64_", `[12]`},
			{125, "repeated_bool_", `[true]`},
			{127, "repeated_string_", `["14"]`},
			{129, "repeated_bytes_", `[` + b64bytes(src.RepeatedBytes_[0]) + `]`},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, typeName, rec.TypeName)
		assert.Len(t, rec.Fields, 15)
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			switch exp.Number {
			case 101, 103:
				assert.InDeltaSlicef(t,
					jsonNumbersAsFloat64(exp.JSON), jsonNumbersAsFloat64(string(field.Value)), epsilon,
					"field=%q", exp.Name)
			default:
				assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
			}
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.RepeatedTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
	t.Run("map", func(t *testing.T) {
		const typeName = "mtechnavi.testdata.MapTypes"
		var src testdatapb.MapTypes
		src.MapStringDouble_ = map[string]float64{"string1": 1.1}
		src.MapStringFloat_ = map[string]float32{"string2": 2.1}
		src.MapStringInt32_ = map[string]int32{"string3": 3}
		src.MapStringInt64_ = map[string]int64{"string4": 4}
		src.MapStringUint32_ = map[string]uint32{"string5": 5}
		src.MapStringUint64_ = map[string]uint64{"string6": 6}
		src.MapStringSint32_ = map[string]int32{"string7": 7}
		src.MapStringSint64_ = map[string]int64{"string8": 8}
		src.MapStringFixed32_ = map[string]uint32{"string9": 9}
		src.MapStringFixed64_ = map[string]uint64{"string10": 10}
		src.MapStringSfixed32_ = map[string]int32{"string11": 11}
		src.MapStringSfixed64_ = map[string]int64{"string12": 12}
		src.MapStringBool_ = map[string]bool{"string13": true}
		src.MapStringString_ = map[string]string{"string14": "14"}
		src.MapStringBytes_ = map[string][]byte{"string15": {0xca, 0xfe, 0xba, 0xbe}}
		src.MapStringMessage_ = map[string]*testdatapb.RootMessage{"string16": {X: "hello"}}
		src.NilMapStringString_ = map[string]string(nil)

		src.MapInt32Double_ = map[int32]float64{1: 1.1}
		src.MapInt32Float_ = map[int32]float32{2: 2.1}
		src.MapInt32Int32_ = map[int32]int32{3: 3}
		src.MapInt32Int64_ = map[int32]int64{4: 4}
		src.MapInt32Uint32_ = map[int32]uint32{5: 5}
		src.MapInt32Uint64_ = map[int32]uint64{6: 6}
		src.MapInt32Sint32_ = map[int32]int32{7: 7}
		src.MapInt32Sint64_ = map[int32]int64{8: 8}
		src.MapInt32Fixed32_ = map[int32]uint32{9: 9}
		src.MapInt32Fixed64_ = map[int32]uint64{10: 10}
		src.MapInt32Sfixed32_ = map[int32]int32{11: 11}
		src.MapInt32Sfixed64_ = map[int32]int64{12: 12}
		src.MapInt32Bool_ = map[int32]bool{13: true}
		src.MapInt32String_ = map[int32]string{14: "14"}
		src.MapInt32Bytes_ = map[int32][]byte{15: {0xca, 0xfe, 0xba, 0xbe}}
		src.MapInt32Message_ = map[int32]*testdatapb.RootMessage{16: {X: "hello"}}
		src.NilMapInt32String_ = map[int32]string(nil)

		src.MapInt64Double_ = map[int64]float64{1: 1.1}
		src.MapInt64Float_ = map[int64]float32{2: 2.1}
		src.MapInt64Int32_ = map[int64]int32{3: 3}
		src.MapInt64Int64_ = map[int64]int64{4: 4}
		src.MapInt64Uint32_ = map[int64]uint32{5: 5}
		src.MapInt64Uint64_ = map[int64]uint64{6: 6}
		src.MapInt64Sint32_ = map[int64]int32{7: 7}
		src.MapInt64Sint64_ = map[int64]int64{8: 8}
		src.MapInt64Fixed32_ = map[int64]uint32{9: 9}
		src.MapInt64Fixed64_ = map[int64]uint64{10: 10}
		src.MapInt64Sfixed32_ = map[int64]int32{11: 11}
		src.MapInt64Sfixed64_ = map[int64]int64{12: 12}
		src.MapInt64Bool_ = map[int64]bool{13: true}
		src.MapInt64String_ = map[int64]string{14: "14"}
		src.MapInt64Bytes_ = map[int64][]byte{15: {0xca, 0xfe, 0xba, 0xbe}}
		src.MapInt64Message_ = map[int64]*testdatapb.RootMessage{16: {X: "hello"}}

		src.MapUint32Double_ = map[uint32]float64{1: 1.1}
		src.MapUint32Float_ = map[uint32]float32{2: 2.1}
		src.MapUint32Int32_ = map[uint32]int32{3: 3}
		src.MapUint32Int64_ = map[uint32]int64{4: 4}
		src.MapUint32Uint32_ = map[uint32]uint32{5: 5}
		src.MapUint32Uint64_ = map[uint32]uint64{6: 6}
		src.MapUint32Sint32_ = map[uint32]int32{7: 7}
		src.MapUint32Sint64_ = map[uint32]int64{8: 8}
		src.MapUint32Fixed32_ = map[uint32]uint32{9: 9}
		src.MapUint32Fixed64_ = map[uint32]uint64{10: 10}
		src.MapUint32Sfixed32_ = map[uint32]int32{11: 11}
		src.MapUint32Sfixed64_ = map[uint32]int64{12: 12}
		src.MapUint32Bool_ = map[uint32]bool{13: true}
		src.MapUint32String_ = map[uint32]string{14: "14"}
		src.MapUint32Bytes_ = map[uint32][]byte{15: {0xca, 0xfe, 0xba, 0xbe}}
		src.MapUint32Message_ = map[uint32]*testdatapb.RootMessage{16: {X: "hello"}}

		src.MapUint64Double_ = map[uint64]float64{1: 1.1}
		src.MapUint64Float_ = map[uint64]float32{2: 2.1}
		src.MapUint64Int32_ = map[uint64]int32{3: 3}
		src.MapUint64Int64_ = map[uint64]int64{4: 4}
		src.MapUint64Uint32_ = map[uint64]uint32{5: 5}
		src.MapUint64Uint64_ = map[uint64]uint64{6: 6}
		src.MapUint64Sint32_ = map[uint64]int32{7: 7}
		src.MapUint64Sint64_ = map[uint64]int64{8: 8}
		src.MapUint64Fixed32_ = map[uint64]uint32{9: 9}
		src.MapUint64Fixed64_ = map[uint64]uint64{10: 10}
		src.MapUint64Sfixed32_ = map[uint64]int32{11: 11}
		src.MapUint64Sfixed64_ = map[uint64]int64{12: 12}
		src.MapUint64Bool_ = map[uint64]bool{13: true}
		src.MapUint64String_ = map[uint64]string{14: "14"}
		src.MapUint64Bytes_ = map[uint64][]byte{15: {0xca, 0xfe, 0xba, 0xbe}}
		src.MapUint64Message_ = map[uint64]*testdatapb.RootMessage{16: {X: "hello"}}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{201, "map_string_double_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string1\"}"`) + `: 1.1}`)},
			{202, "map_string_float_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string2\"}"`) + `: 2.1}`)},
			{203, "map_string_int32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string3\"}"`) + `: 3}`)},
			{204, "map_string_int64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string4\"}"`) + `: 4}`)},
			{205, "map_string_uint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string5\"}"`) + `: 5}`)},
			{206, "map_string_uint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string6\"}"`) + `: 6}`)},
			{207, "map_string_sint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string7\"}"`) + `: 7}`)},
			{208, "map_string_sint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string8\"}"`) + `: 8}`)},
			{209, "map_string_fixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string9\"}"`) + `: 9}`)},
			{210, "map_string_fixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string10\"}"`) + `: 10}`)},
			{211, "map_string_sfixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string11\"}"`) + `: 11}`)},
			{212, "map_string_sfixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string12\"}"`) + `: 12}`)},
			{213, "map_string_bool_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string13\"}"`) + `: true}`)},
			{214, "map_string_string_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string14\"}"`) + `: "14"}`)},
			{215, "map_string_bytes_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string15\"}"`) + `: ` + b64bytes(src.MapStringBytes_["string15"]) + `}`)},
			{216, "map_string_message_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"string\",\"value\":\"string16\"}"`) + `: ` +
				jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields": {
						"1":{"name":"x","value":`+b64bytes(jsonBytes("hello"))+`,"visibility":{}},
						"2":{"name":"shared_properties","value":`+b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{})))+`,"visibility":{}},
						"3":{"name":"created_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"4":{"name":"updated_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"5":{"name":"deleted_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}}
					}
				}`) + `}`),
			},
			{217, "nil_map_string_string_", "{}"},

			{221, "map_int32_double_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":1}"`) + `: 1.1}`)},
			{222, "map_int32_float_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":2}"`) + `: 2.1}`)},
			{223, "map_int32_int32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":3}"`) + `: 3}`)},
			{224, "map_int32_int64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":4}"`) + `: 4}`)},
			{225, "map_int32_uint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":5}"`) + `: 5}`)},
			{226, "map_int32_uint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":6}"`) + `: 6}`)},
			{227, "map_int32_sint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":7}"`) + `: 7}`)},
			{228, "map_int32_sint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":8}"`) + `: 8}`)},
			{229, "map_int32_fixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":9}"`) + `: 9}`)},
			{230, "map_int32_fixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":10}"`) + `: 10}`)},
			{231, "map_int32_sfixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":11}"`) + `: 11}`)},
			{232, "map_int32_sfixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":12}"`) + `: 12}`)},
			{233, "map_int32_bool_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":13}"`) + `: true}`)},
			{234, "map_int32_string_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":14}"`) + `: "14"}`)},
			{235, "map_int32_bytes_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":15}"`) + `: ` + b64bytes(src.MapStringBytes_["string15"]) + `}`)},
			{236, "map_int32_message_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int32\",\"value\":16}"`) + `: ` +
				jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields": {
						"1":{"name":"x","value":`+b64bytes(jsonBytes("hello"))+`,"visibility":{}},
						"2":{"name":"shared_properties","value":`+b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{})))+`,"visibility":{}},
						"3":{"name":"created_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"4":{"name":"updated_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"5":{"name":"deleted_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}}
					}
				}`) + `}`),
			},
			{237, "nil_map_int32_string_", "{}"},

			{241, "map_int64_double_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":1}"`) + `: 1.1}`)},
			{242, "map_int64_float_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":2}"`) + `: 2.1}`)},
			{243, "map_int64_int32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":3}"`) + `: 3}`)},
			{244, "map_int64_int64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":4}"`) + `: 4}`)},
			{245, "map_int64_uint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":5}"`) + `: 5}`)},
			{246, "map_int64_uint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":6}"`) + `: 6}`)},
			{247, "map_int64_sint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":7}"`) + `: 7}`)},
			{248, "map_int64_sint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":8}"`) + `: 8}`)},
			{249, "map_int64_fixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":9}"`) + `: 9}`)},
			{250, "map_int64_fixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":10}"`) + `: 10}`)},
			{251, "map_int64_sfixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":11}"`) + `: 11}`)},
			{252, "map_int64_sfixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":12}"`) + `: 12}`)},
			{253, "map_int64_bool_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":13}"`) + `: true}`)},
			{254, "map_int64_string_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":14}"`) + `: "14"}`)},
			{255, "map_int64_bytes_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":15}"`) + `: ` + b64bytes(src.MapStringBytes_["string15"]) + `}`)},
			{256, "map_int64_message_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"int64\",\"value\":16}"`) + `: ` +
				jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields": {
						"1":{"name":"x","value":`+b64bytes(jsonBytes("hello"))+`,"visibility":{}},
						"2":{"name":"shared_properties","value":`+b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{})))+`,"visibility":{}},
						"3":{"name":"created_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"4":{"name":"updated_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"5":{"name":"deleted_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}}
					}
				}`) + `}`),
			},
			{257, "nil_map_int64_string_", "{}"},

			{261, "map_uint32_double_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":1}"`) + `: 1.1}`)},
			{262, "map_uint32_float_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":2}"`) + `: 2.1}`)},
			{263, "map_uint32_int32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":3}"`) + `: 3}`)},
			{264, "map_uint32_int64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":4}"`) + `: 4}`)},
			{265, "map_uint32_uint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":5}"`) + `: 5}`)},
			{266, "map_uint32_uint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":6}"`) + `: 6}`)},
			{267, "map_uint32_sint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":7}"`) + `: 7}`)},
			{268, "map_uint32_sint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":8}"`) + `: 8}`)},
			{269, "map_uint32_fixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":9}"`) + `: 9}`)},
			{270, "map_uint32_fixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":10}"`) + `: 10}`)},
			{271, "map_uint32_sfixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":11}"`) + `: 11}`)},
			{272, "map_uint32_sfixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":12}"`) + `: 12}`)},
			{273, "map_uint32_bool_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":13}"`) + `: true}`)},
			{274, "map_uint32_string_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":14}"`) + `: "14"}`)},
			{275, "map_uint32_bytes_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":15}"`) + `: ` + b64bytes(src.MapStringBytes_["string15"]) + `}`)},
			{276, "map_uint32_message_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint32\",\"value\":16}"`) + `: ` +
				jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields": {
						"1":{"name":"x","value":`+b64bytes(jsonBytes("hello"))+`,"visibility":{}},
						"2":{"name":"shared_properties","value":`+b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{})))+`,"visibility":{}},
						"3":{"name":"created_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"4":{"name":"updated_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"5":{"name":"deleted_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}}
					}
				}`) + `}`),
			},
			{277, "nil_map_uint32_string_", "{}"},

			{281, "map_uint64_double_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":1}"`) + `: 1.1}`)},
			{282, "map_uint64_float_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":2}"`) + `: 2.1}`)},
			{283, "map_uint64_int32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":3}"`) + `: 3}`)},
			{284, "map_uint64_int64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":4}"`) + `: 4}`)},
			{285, "map_uint64_uint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":5}"`) + `: 5}`)},
			{286, "map_uint64_uint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":6}"`) + `: 6}`)},
			{287, "map_uint64_sint32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":7}"`) + `: 7}`)},
			{288, "map_uint64_sint64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":8}"`) + `: 8}`)},
			{289, "map_uint64_fixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":9}"`) + `: 9}`)},
			{290, "map_uint64_fixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":10}"`) + `: 10}`)},
			{291, "map_uint64_sfixed32_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":11}"`) + `: 11}`)},
			{292, "map_uint64_sfixed64_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":12}"`) + `: 12}`)},
			{293, "map_uint64_bool_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":13}"`) + `: true}`)},
			{294, "map_uint64_string_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":14}"`) + `: "14"}`)},
			{295, "map_uint64_bytes_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":15}"`) + `: ` + b64bytes(src.MapStringBytes_["string15"]) + `}`)},
			{296, "map_uint64_message_", jsonCompact(`{` + jsonCompact(`"{\"type\":\"uint64\",\"value\":16}"`) + `: ` +
				jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields": {
						"1":{"name":"x","value":`+b64bytes(jsonBytes("hello"))+`,"visibility":{}},
						"2":{"name":"shared_properties","value":`+b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{})))+`,"visibility":{}},
						"3":{"name":"created_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"4":{"name":"updated_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}},
						"5":{"name":"deleted_at","value":`+b64bytes(jsonBytes(int64(0)))+`,"visibility":{}}
					}
				}`) + `}`),
			},
			{297, "nil_map_uint64_string_", "{}"},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, typeName, rec.TypeName)
		assert.Len(t, rec.Fields, 85)
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			switch exp.Number {
			case 201, 202, 221, 222, 241, 242, 261, 262, 281, 282:
				assert.InDeltaMapValuesf(t,
					jsonMapAsFloat64(exp.JSON), jsonMapAsFloat64(string(field.Value)), epsilon,
					"[%d] %s", exp.Number, exp.Name)
			default:
				assert.Equalf(t, exp.JSON, string(field.Value), "[%d] %s", exp.Number, exp.Name)
			}
		}

		// decode from encoded
		var dst testdatapb.MapTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// omit zero value by encode
			src.NilMapInt32String_ = map[int32]string{}
			src.MapInt32Message_[16].SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
			src.NilMapInt64String_ = map[int64]string{}
			src.MapInt64Message_[16].SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
			src.NilMapStringString_ = map[string]string{}
			src.MapStringMessage_["string16"].SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
			src.NilMapUint32String_ = map[uint32]string{}
			src.MapUint32Message_[16].SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
			src.NilMapUint64String_ = map[uint64]string{}
			src.MapUint64Message_[16].SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
}

func TestOneOfTypes(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		const typeName = "mtechnavi.testdata.OneOfTypes"
		var src testdatapb.OneOfTypes

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{3, "message1", jsonCompact(`null`)},
			{5, "message2", jsonCompact(`null`)},
		}
		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, typeName, rec.TypeName)
		assert.Len(t, rec.Fields, 2)
		for _, exp := range expects {
			field, ok := rec.Fields[int32(exp.Number)]
			assert.Truef(t, ok, "field=%q", exp.Name)
			assert.Equal(t, exp.Name, field.Name, "field=%q", exp.Name)
			assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.OneOfTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			// assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
			// assert.JSONEq(t, jsonStr(src.GetName()), jsonStr(dst.GetName()))
			assert.JSONEq(t, jsonStr(src.GetMessage1()), jsonStr(dst.GetMessage1()))
			assert.JSONEq(t, jsonStr(src.GetMessage2()), jsonStr(dst.GetMessage2()))
			// assert.Equal(t, "", dst.GetName())
			assert.Nil(t, dst.GetMessage1())
			assert.Nil(t, dst.GetMessage2())

		}
	})

	t.Run("not zero value", func(t *testing.T) {
		const typeName = "mtechnavi.testdata.OneOfTypes"
		var src testdatapb.OneOfTypes
		message1 := &testdatapb.RootMessage{
			X:                "a1",
			SharedProperties: &sharelibpb.EmbeddedSharedProperties{},
		}
		src.TestOneof = &testdatapb.OneOfTypes_Message1{
			Message1: message1,
		}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{3, "message1", jsonCompact(`{
					"type_name":"mtechnavi.testdata.RootMessage",
					"fields":{
						"1":{"name":"x","value":` + b64bytes(jsonBytes(string("a1"))) + `,"visibility":{}},
						"2":{"name":"shared_properties","value":` + b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{}))) + `,"visibility":{}},
						"3":{"name":"created_at","value":` + b64bytes(jsonBytes(int64(0))) + `,"visibility":{}},
						"4":{"name":"updated_at","value":` + b64bytes(jsonBytes(int64(0))) + `,"visibility":{}},
						"5":{"name":"deleted_at","value":` + b64bytes(jsonBytes(int64(0))) + `,"visibility":{}}
					},
					"shared_properties":{}
				}`)},
			{5, "message2", jsonCompact(`null`)},
		}
		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, typeName, rec.TypeName)
		assert.Len(t, rec.Fields, 2)
		for _, exp := range expects {
			field, ok := rec.Fields[int32(exp.Number)]
			assert.Truef(t, ok, "field=%q", exp.Name)
			assert.Equal(t, exp.Name, field.Name, "field=%q", exp.Name)
			assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.OneOfTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			// assert.Equal(t, jsonStr(&src), jsonStr(&dst))
			// assert.JSONEq(t, jsonStr(src.GetName()), jsonStr(dst.GetName()))
			assert.JSONEq(t, jsonStr(src.GetMessage1()), jsonStr(dst.GetMessage1()))
			assert.JSONEq(t, jsonStr(src.GetMessage2()), jsonStr(dst.GetMessage2()))
			// assert.Equal(t, "", dst.GetName())
			assert.NotNil(t, dst.GetMessage1())
			assert.NotNil(t, dst.GetTestOneof())
			assert.Nil(t, dst.GetMessage2())
		}
	})
}

func TestMessageTypes(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		var v testdatapb.MessageTypes

		expect := map[int32]*recordpb.Field{
			1: {
				Name:       "root_message1",
				Value:      jsonBytes(encode(&testdatapb.RootMessage{})),
				Visibility: &mtn.VisibilityOptions{},
			},
			2: {
				Name:       "root_message2",
				Value:      jsonBytes([]interface{}{}),
				Visibility: &mtn.VisibilityOptions{},
			},
			3: {
				Name:       "scoped_message1",
				Value:      jsonBytes(encode(&testdatapb.MessageTypes_ScopedMessage{})),
				Visibility: &mtn.VisibilityOptions{},
			},
			4: {
				Name:       "scoped_message2",
				Value:      jsonBytes([]interface{}{}),
				Visibility: &mtn.VisibilityOptions{},
			},
		}
		rec, err := codec.Encode(&v)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "mtechnavi.testdata.MessageTypes", rec.TypeName)
		assert.Len(t, rec.Fields, len(expect))
		assert.Equalf(t, expect, rec.Fields, "field=%s", "Fields")
	})
	t.Run("single", func(t *testing.T) {
		var src testdatapb.MessageTypes
		src.RootMessage1 = &testdatapb.RootMessage{
			X: "hello",
			SharedProperties: &sharelibpb.EmbeddedSharedProperties{
				SharedBy: "tenant001",
				SharedAt: 123,
			},
			CreatedAt: 4,
			UpdatedAt: 5,
			DeletedAt: 6,
		}
		src.ScopedMessage1 = &testdatapb.MessageTypes_ScopedMessage{
			Y: "world",
		}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{1, "root_message1", jsonCompact(`{
				"type_name": "mtechnavi.testdata.RootMessage",
				"fields": {
					"1": {
						"name":"x",
						"value":` + b64bytes(jsonBytes("hello")) + `,
						"visibility":{}
					},
					"2":{
						"name":"shared_properties",
						"value":` + b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{
				SharedBy: "tenant001",
				SharedAt: 123,
			}))) + `,
						"visibility":{}
					},
					"3":{
						"name":"created_at",
						"value":` + b64bytes(jsonBytes(int64(4))) + `,
						"visibility":{}
					},
					"4":{
						"name":"updated_at",
						"value":` + b64bytes(jsonBytes(int64(5))) + `,
						"visibility":{}
					},
					"5":{
						"name":"deleted_at",
						"value":` + b64bytes(jsonBytes(int64(6))) + `,
						"visibility":{}
					}
				},
				"created_at":4,
				"updated_at":5,
				"deleted_at":6,
				"shared_properties":{
					"shared_by":"tenant001",
					"shared_at":123
				}
			}`)},
			{3, "scoped_message1", jsonCompact(`{
				"type_name": "mtechnavi.testdata.MessageTypes.ScopedMessage",
				"fields": {
					"1": {
						"name":"y",
						"value":` + b64bytes(jsonBytes("world")) + `,
						"visibility":{}
					}
				}
			}`)},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "mtechnavi.testdata.MessageTypes", rec.TypeName)
		assert.Len(t, rec.Fields, 4)
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			assert.JSONEqf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.MessageTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
	t.Run("repeated", func(t *testing.T) {
		var src testdatapb.MessageTypes
		src.RootMessage1 = &testdatapb.RootMessage{
			SharedProperties: &sharelibpb.EmbeddedSharedProperties{},
		}
		src.RootMessage2 = []*testdatapb.RootMessage{
			{
				X:                "hello",
				SharedProperties: &sharelibpb.EmbeddedSharedProperties{},
			},
		}
		src.ScopedMessage1 = &testdatapb.MessageTypes_ScopedMessage{}
		src.ScopedMessage2 = []*testdatapb.MessageTypes_ScopedMessage{
			{
				Y: "world",
			},
		}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{2, "root_message2", jsonCompact(`[{
				"type_name":"mtechnavi.testdata.RootMessage",
				"fields":{
					"1":{
						"name":"x",
						"value":` + b64bytes(jsonBytes("hello")) + `,
						"visibility":{}
					},
					"2":{
						"name":"shared_properties",
						"value":` + b64bytes(jsonBytes(encode(&sharelibpb.EmbeddedSharedProperties{}))) + `,
						"visibility":{}
					},
					"3":{
						"name":"created_at",
						"value":` + b64bytes(jsonBytes(int64(0))) + `,
						"visibility":{}
					},
					"4":{
						"name":"updated_at",
						"value":` + b64bytes(jsonBytes(int64(0))) + `,
						"visibility":{}
					},
					"5":{
						"name":"deleted_at",
						"value":` + b64bytes(jsonBytes(int64(0))) + `,
						"visibility":{}
					}
				},
				"shared_properties":{}
			}]`)},
			{4, "scoped_message2", jsonCompact(`[{
				"type_name": "mtechnavi.testdata.MessageTypes.ScopedMessage",
				"fields": {
					"1": {
						"name":"y",
						"value":` + b64bytes(jsonBytes("world")) + `,
						"visibility":{}
					}
				}
			}]`)},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "mtechnavi.testdata.MessageTypes", rec.TypeName)
		assert.Len(t, rec.Fields, 4)
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.MessageTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
}

func TestEnumTypes(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		var v testdatapb.EnumTypes
		expect := map[int32]*recordpb.Field{
			1: {
				Name:       "root_enum1",
				Value:      jsonBytes(int32(0)),
				Visibility: &mtn.VisibilityOptions{},
			},
			3: {
				Name:       "root_enum2",
				Value:      jsonBytes(make([]interface{}, 0)),
				Visibility: &mtn.VisibilityOptions{},
			},
			5: {
				Name:       "scoped_enum1",
				Value:      jsonBytes(int32(0)),
				Visibility: &mtn.VisibilityOptions{},
			},
			7: {
				Name:       "scoped_enum2",
				Value:      jsonBytes(make([]interface{}, 0)),
				Visibility: &mtn.VisibilityOptions{},
			},
		}
		rec, err := codec.Encode(&v)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "mtechnavi.testdata.EnumTypes", rec.TypeName)
		assert.Len(t, rec.Fields, len(expect))
		assert.Equalf(t, expect, rec.Fields, "field=%s", "Fields")
	})
	t.Run("single", func(t *testing.T) {
		var src testdatapb.EnumTypes
		src.RootEnum1 = testdatapb.RootEnum_B
		src.ScopedEnum1 = testdatapb.EnumTypes_Y
		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "mtechnavi.testdata.EnumTypes", rec.TypeName)
		assert.Len(t, rec.Fields, 4)
		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{1, "root_enum1", `1`},
			{5, "scoped_enum1", `1`},
		}
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.EnumTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
	t.Run("repeated", func(t *testing.T) {
		var src testdatapb.EnumTypes
		src.RootEnum2 = []testdatapb.RootEnum{
			testdatapb.RootEnum_A,
			testdatapb.RootEnum_B,
			testdatapb.RootEnum_C,
		}
		src.ScopedEnum2 = []testdatapb.EnumTypes_ScopedEnum{
			testdatapb.EnumTypes_X,
			testdatapb.EnumTypes_Y,
			testdatapb.EnumTypes_Z,
		}

		expects := []struct {
			Number int
			Name   string
			JSON   string
		}{
			{3, "root_enum2", jsonCompact(`[0,1,2]`)},
			{7, "scoped_enum2", jsonCompact(`[0,1,2]`)},
		}

		rec, err := codec.Encode(&src)
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, "mtechnavi.testdata.EnumTypes", rec.TypeName)
		assert.Len(t, rec.Fields, 4)
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			assert.Equalf(t, exp.JSON, string(field.Value), "field=%q", exp.Name)
		}

		// codec.Debug(os.Stderr, rec)

		// decode from encoded
		var dst testdatapb.EnumTypes
		if assert.NoError(t, codec.Decode(&dst, rec)) {
			// 内部フィールドなどは比較したくないため、一度JSON化して比較する
			assert.JSONEq(t, jsonStr(&src), jsonStr(&dst))
		}
	})
}

func TestVisibilityTypes(t *testing.T) {
	t.Run("visibility options", func(t *testing.T) {
		var v = testdatapb.VisibilityTypes{
			Private1: "private1",
			Private2: "private2",
			Granted:  "granted",
			Public:   "public",
		}
		rec, err := codec.Encode(&v)
		if !assert.NoError(t, err) {
			return
		}
		expects := []struct {
			Number     int
			Name       string
			Visibility *mtn.VisibilityOptions
		}{
			{1, "private1", &mtn.VisibilityOptions{Scope: mtn.Scope_PRIVATE}},
			{2, "private2", &mtn.VisibilityOptions{Scope: mtn.Scope_PRIVATE}},
			{3, "granted", &mtn.VisibilityOptions{Scope: mtn.Scope_GRANTED}},
			{4, "public", &mtn.VisibilityOptions{Scope: mtn.Scope_PUBLIC}},
		}
		for _, exp := range expects {
			field := rec.Fields[int32(exp.Number)]
			assert.Equal(t, exp.Name, field.Name)
			assert.Equal(t, exp.Visibility.Scope, field.Visibility.Scope, "field=%q", exp.Name)
		}
		assert.Equal(t, "mtechnavi.testdata.VisibilityTypes", rec.TypeName)
	})
}

func jsonCompact(s string) string {
	var dst bytes.Buffer
	if err := json.Compact(&dst, []byte(s)); err != nil {
		panic(err)
	}
	return dst.String()
}

func jsonStr(v interface{}) string {
	return string(jsonBytes(v))
}

func jsonBytes(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func b64bytes(v []byte) string {
	return `"` + base64.StdEncoding.EncodeToString(v) + `"`
}

func jsonNumberAsFloat64(data string) float64 {
	v, err := json.Number(data).Float64()
	if err != nil {
		panic(fmt.Sprintf("invalid json: %q: %v", data, err))
	}
	return v
}

func jsonNumbersAsFloat64(data string) []float64 {
	var l []json.Number
	if err := json.Unmarshal([]byte(data), &l); err != nil {
		panic(fmt.Sprintf("%v: %q", err, data))
	}
	out := make([]float64, len(l))
	for i := range l {
		out[i] = jsonNumberAsFloat64(l[i].String())
	}
	return out
}

func jsonMapAsFloat64(data string) map[string]float64 {
	var m map[string]json.Number
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}
	out := map[string]float64{}
	for k, v := range m {
		out[k] = jsonNumberAsFloat64(v.String())
	}
	return out
}

func jsonPP(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func encode(v proto.Message) *recordpb.Record {
	r, err := codec.Encode(v)
	if err != nil {
		panic(err)
	}
	return r
}
