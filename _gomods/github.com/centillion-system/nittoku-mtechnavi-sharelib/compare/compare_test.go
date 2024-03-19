package compare_test

import (
	"encoding/json"
	"mtechnavi/sharelib/compare"
	testdatapb "mtechnavi/sharelib/compare/testdata/protobuf"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestDiff(t *testing.T) {
	var noDiff []*compare.Diff
	type args struct {
		scope mtnpb.Scope
		x     proto.Message
		y     proto.Message
	}
	tests := []struct {
		name    string
		args    args
		want    []*compare.Diff
		wantErr bool
	}{
		{
			name: "scalar_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.ScalarTypes{},
				y:     &testdatapb.ScalarTypes{},
			},
			want: noDiff,
		},
		{
			name: "scalar_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.ScalarTypes{
					PrivateDouble_:   1.1,
					GrantedFloat_:    2.1,
					PublicInt32_:     3,
					PrivateInt64_:    4,
					GrantedUint32_:   5,
					PublicUint64_:    6,
					PrivateSint32_:   7,
					GrantedSint64_:   8,
					PublicFixed32_:   9,
					PrivateFixed64_:  10,
					GrantedSfixed32_: 11,
					PublicSfixed64_:  12,
					PrivateBool_:     true,
					GrantedString_:   "14",
					PublicBytes_:     []byte("hello"),
				},
				y: &testdatapb.ScalarTypes{
					PrivateDouble_:   1.12,
					GrantedFloat_:    2.12,
					PublicInt32_:     32,
					PrivateInt64_:    42,
					GrantedUint32_:   52,
					PublicUint64_:    62,
					PrivateSint32_:   72,
					GrantedSint64_:   82,
					PublicFixed32_:   92,
					PrivateFixed64_:  102,
					GrantedSfixed32_: 112,
					PublicSfixed64_:  122,
					PrivateBool_:     false,
					GrantedString_:   "142",
					PublicBytes_:     []byte("world"),
				},
			},
			want: []*compare.Diff{
				{Path: "privateDouble", ValueX: float64(1.1), ValueY: float64(1.12)},
				{Path: "grantedFloat", ValueX: float32(2.1), ValueY: float32(2.12)},
				{Path: "publicInt32", ValueX: int32(3), ValueY: int32(32)},
				{Path: "privateInt64", ValueX: int64(4), ValueY: int64(42)},
				{Path: "grantedUint32", ValueX: uint32(5), ValueY: uint32(52)},
				{Path: "publicUint64", ValueX: uint64(6), ValueY: uint64(62)},
				{Path: "privateSint32", ValueX: int32(7), ValueY: int32(72)},
				{Path: "grantedSint64", ValueX: int64(8), ValueY: int64(82)},
				{Path: "publicFixed32", ValueX: uint32(9), ValueY: uint32(92)},
				{Path: "privateFixed64", ValueX: uint64(10), ValueY: uint64(102)},
				{Path: "grantedSfixed32", ValueX: int32(11), ValueY: int32(112)},
				{Path: "publicSfixed64", ValueX: int64(12), ValueY: int64(122)},
				{Path: "privateBool", ValueX: bool(true), ValueY: bool(false)},
				{Path: "grantedString", ValueX: "14", ValueY: "142"},
				{Path: "publicBytes", ValueX: []byte("hello"), ValueY: []byte("world")},
			},
		},
		{
			name: "repeated_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.RepeatedTypes{},
				y:     &testdatapb.RepeatedTypes{},
			},
			want: noDiff,
		},
		{
			name: "repeated_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.RepeatedTypes{
					PrivateRepeatedDouble_:   []float64{1.1, 1.2, 1.3},
					GrantedRepeatedFloat_:    []float32{2.1, 2.2, 2.3},
					PublicRepeatedInt32_:     []int32{31, 32, 33},
					PrivateRepeatedInt64_:    []int64{41, 42, 43},
					GrantedRepeatedUint32_:   []uint32{51, 52, 53},
					PublicRepeatedUint64_:    []uint64{61, 62, 63},
					PrivateRepeatedSint32_:   []int32{71, 72, 73},
					GrantedRepeatedSint64_:   []int64{81, 82, 83},
					PublicRepeatedFixed32_:   []uint32{91, 92, 93},
					PrivateRepeatedFixed64_:  []uint64{101, 102, 103},
					GrantedRepeatedSfixed32_: []int32{111, 112, 113},
					PublicRepeatedSfixed64_:  []int64{121, 122, 123},
					PrivateRepeatedBool_:     []bool{true, false, true},
					GrantedRepeatedString_:   []string{"141", "142", "143"},
					PublicRepeatedBytes_:     [][]byte{[]byte("this"), []byte("is"), []byte("test")},
					PrivateRepeatedMessage_: []*testdatapb.ScalarTypes{
						{PrivateDouble_: 1.1, GrantedFloat_: 2.1, PublicInt32_: 31, GrantedString_: "valueX_GrantedString_1"},
						{PrivateDouble_: 1.2, GrantedFloat_: 2.2, PublicInt32_: 32},
						{PrivateDouble_: 1.3, GrantedFloat_: 2.3, PublicInt32_: 33},
					},
				},
				y: &testdatapb.RepeatedTypes{
					PrivateRepeatedDouble_:   []float64{1.12, 1.22, 1.32},
					GrantedRepeatedFloat_:    []float32{2.12, 2.22, 2.32},
					PublicRepeatedInt32_:     []int32{312, 322, 332},
					PrivateRepeatedInt64_:    []int64{412, 422, 432},
					GrantedRepeatedUint32_:   []uint32{512, 522, 532},
					PublicRepeatedUint64_:    []uint64{612, 622, 632},
					PrivateRepeatedSint32_:   []int32{712, 722, 732},
					GrantedRepeatedSint64_:   []int64{812, 822, 832},
					PublicRepeatedFixed32_:   []uint32{912, 922, 932},
					PrivateRepeatedFixed64_:  []uint64{1012, 1022, 1032},
					GrantedRepeatedSfixed32_: []int32{1112, 1122, 1132},
					PublicRepeatedSfixed64_:  []int64{1212, 1222, 1232},
					PrivateRepeatedBool_:     []bool{false, true, false},
					GrantedRepeatedString_:   []string{"1412", "1422", "1432"},
					PublicRepeatedBytes_:     [][]byte{[]byte("test"), []byte("is"), []byte("this")},
					PrivateRepeatedMessage_: []*testdatapb.ScalarTypes{
						{PrivateDouble_: 1.12, GrantedFloat_: 2.12, PublicInt32_: 312},
						{PrivateDouble_: 1.22, GrantedFloat_: 2.22, PublicInt32_: 322},
						{PrivateDouble_: 1.32, GrantedFloat_: 2.32, PublicInt32_: 332},
					},
				},
			},
			want: []*compare.Diff{
				{Path: "privateRepeatedDouble", ValueX: []float64{1.1, 1.2, 1.3}, ValueY: []float64{1.12, 1.22, 1.32}},
				{Path: "grantedRepeatedFloat", ValueX: []float32{2.1, 2.2, 2.3}, ValueY: []float32{2.12, 2.22, 2.32}},
				{Path: "publicRepeatedInt32", ValueX: []int32{31, 32, 33}, ValueY: []int32{312, 322, 332}},
				{Path: "privateRepeatedInt64", ValueX: []int64{41, 42, 43}, ValueY: []int64{412, 422, 432}},
				{Path: "grantedRepeatedUint32", ValueX: []uint32{51, 52, 53}, ValueY: []uint32{512, 522, 532}},
				{Path: "publicRepeatedUint64", ValueX: []uint64{61, 62, 63}, ValueY: []uint64{612, 622, 632}},
				{Path: "privateRepeatedSint32", ValueX: []int32{71, 72, 73}, ValueY: []int32{712, 722, 732}},
				{Path: "grantedRepeatedSint64", ValueX: []int64{81, 82, 83}, ValueY: []int64{812, 822, 832}},
				{Path: "publicRepeatedFixed32", ValueX: []uint32{91, 92, 93}, ValueY: []uint32{912, 922, 932}},
				{Path: "privateRepeatedFixed64", ValueX: []uint64{101, 102, 103}, ValueY: []uint64{1012, 1022, 1032}},
				{Path: "grantedRepeatedSfixed32", ValueX: []int32{111, 112, 113}, ValueY: []int32{1112, 1122, 1132}},
				{Path: "publicRepeatedSfixed64", ValueX: []int64{121, 122, 123}, ValueY: []int64{1212, 1222, 1232}},
				{Path: "privateRepeatedBool", ValueX: []bool{true, false, true}, ValueY: []bool{false, true, false}},
				{Path: "grantedRepeatedString", ValueX: []string{"141", "142", "143"}, ValueY: []string{"1412", "1422", "1432"}},
				{Path: "publicRepeatedBytes", ValueX: [][]byte{[]byte("this"), []byte("is"), []byte("test")}, ValueY: [][]byte{[]byte("test"), []byte("is"), []byte("this")}},
				{Path: "privateRepeatedMessage",
					ValueX: []*testdatapb.ScalarTypes{
						{PrivateDouble_: 1.1, GrantedFloat_: 2.1, PublicInt32_: 31, GrantedString_: "valueX_GrantedString_1"},
						{PrivateDouble_: 1.2, GrantedFloat_: 2.2, PublicInt32_: 32},
						{PrivateDouble_: 1.3, GrantedFloat_: 2.3, PublicInt32_: 33},
					},
					ValueY: []*testdatapb.ScalarTypes{
						{PrivateDouble_: 1.12, GrantedFloat_: 2.12, PublicInt32_: 312},
						{PrivateDouble_: 1.22, GrantedFloat_: 2.22, PublicInt32_: 322},
						{PrivateDouble_: 1.32, GrantedFloat_: 2.32, PublicInt32_: 332},
					},
				},
			},
		},
		{
			name: "map_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.MapTypes{},
				y:     &testdatapb.MapTypes{},
			},
			want: noDiff,
		},
		{
			name: "map_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.MapTypes{
					PrivateMapStringDouble_:   map[string]float64{"PrivateMapStringDouble_1": float64(1.1), "PrivateMapStringDouble_2": float64(1.2), "PrivateMapStringDouble_3": float64(1.3)},
					GrantedMapStringFloat_:    map[string]float32{"GrantedMapStringFloat_1": float32(2.1), "GrantedMapStringFloat_2": float32(2.2), "GrantedMapStringFloat_3": float32(2.3)},
					PublicMapStringInt32_:     map[string]int32{"PublicMapStringInt32_1": int32(31), "PublicMapStringInt32_2": int32(32), "PublicMapStringInt32_3": int32(33)},
					PrivateMapStringInt64_:    map[string]int64{"PrivateMapStringInt64_1": int64(41), "PrivateMapStringInt64_2": int64(42), "PrivateMapStringInt64_3": int64(43)},
					GrantedMapStringUint32_:   map[string]uint32{"GrantedMapStringUint32_1": uint32(51), "GrantedMapStringUint32_2": uint32(52), "GrantedMapStringUint32_3": uint32(53)},
					PublicMapStringUint64_:    map[string]uint64{"PublicMapStringUint64_1": uint64(61), "PublicMapStringUint64_2": uint64(62), "PublicMapStringUint64_3": uint64(63)},
					PrivateMapStringSint32_:   map[string]int32{"PrivateMapStringSint32_1": int32(71), "PrivateMapStringSint32_2": int32(72), "PrivateMapStringSint32_3": int32(73)},
					GrantedMapStringSint64_:   map[string]int64{"GrantedMapStringSint64_1": int64(81), "GrantedMapStringSint64_2": int64(82), "GrantedMapStringSint64_3": int64(83)},
					PublicMapStringFixed32_:   map[string]uint32{"PublicMapStringFixed32_1": uint32(91), "PublicMapStringFixed32_2": uint32(92), "PublicMapStringFixed32_3": uint32(93)},
					PrivateMapStringFixed64_:  map[string]uint64{"PrivateMapStringFixed64_1": uint64(101), "PrivateMapStringFixed64_2": uint64(102), "PrivateMapStringFixed64_3": uint64(103)},
					GrantedMapStringSfixed32_: map[string]int32{"GrantedMapStringSfixed32_1": int32(111), "GrantedMapStringSfixed32_2": int32(112), "GrantedMapStringSfixed32_3": int32(113)},
					PublicMapStringSfixed64_:  map[string]int64{"PublicMapStringSfixed64_1": int64(121), "PublicMapStringSfixed64_2": int64(122), "PublicMapStringSfixed64_3": int64(123)},
					PrivateMapStringBool_:     map[string]bool{"PrivateMapStringBool_1": bool(true), "PrivateMapStringBool_2": bool(false), "PrivateMapStringBool_3": bool(true)},
					GrantedMapStringString_:   map[string]string{"GrantedMapStringString_1": string("141"), "GrantedMapStringString_2": string("142"), "GrantedMapStringString_3": string("143")},
					PublicMapStringBytes_:     map[string][]byte{"PublicMapStringBytes_1": []byte("this"), "PublicMapStringBytes_2": []byte("is"), "PublicMapStringBytes_3": []byte("test")},
					PrivateMapStringMessage_: map[string]*testdatapb.ScalarTypes{
						"PrivateMapStringMessage_1": {PrivateDouble_: 1.1, GrantedFloat_: 2.1, PublicInt32_: 31, GrantedString_: "valueX_GrantedString_1"},
						"PrivateMapStringMessage_2": {PrivateDouble_: 1.2, GrantedFloat_: 2.2, PublicInt32_: 32},
						"PrivateMapStringMessage_3": {PrivateDouble_: 1.3, GrantedFloat_: 2.3, PublicInt32_: 33},
					},
				},
				y: &testdatapb.MapTypes{
					PrivateMapStringDouble_:   map[string]float64{"PrivateMapStringDouble_1": float64(1.12), "PrivateMapStringDouble_2": float64(1.22), "PrivateMapStringDouble_3": float64(1.32)},
					GrantedMapStringFloat_:    map[string]float32{"GrantedMapStringFloat_1": float32(2.12), "GrantedMapStringFloat_2": float32(2.22), "GrantedMapStringFloat_3": float32(2.32)},
					PublicMapStringInt32_:     map[string]int32{"PublicMapStringInt32_1": int32(312), "PublicMapStringInt32_2": int32(322), "PublicMapStringInt32_3": int32(332)},
					PrivateMapStringInt64_:    map[string]int64{"PrivateMapStringInt64_1": int64(412), "PrivateMapStringInt64_2": int64(422), "PrivateMapStringInt64_3": int64(432)},
					GrantedMapStringUint32_:   map[string]uint32{"GrantedMapStringUint32_1": uint32(512), "GrantedMapStringUint32_2": uint32(522), "GrantedMapStringUint32_3": uint32(532)},
					PublicMapStringUint64_:    map[string]uint64{"PublicMapStringUint64_1": uint64(612), "PublicMapStringUint64_2": uint64(622), "PublicMapStringUint64_3": uint64(632)},
					PrivateMapStringSint32_:   map[string]int32{"PrivateMapStringSint32_1": int32(712), "PrivateMapStringSint32_2": int32(722), "PrivateMapStringSint32_3": int32(732)},
					GrantedMapStringSint64_:   map[string]int64{"GrantedMapStringSint64_1": int64(812), "GrantedMapStringSint64_2": int64(822), "GrantedMapStringSint64_3": int64(832)},
					PublicMapStringFixed32_:   map[string]uint32{"PublicMapStringFixed32_1": uint32(912), "PublicMapStringFixed32_2": uint32(922), "PublicMapStringFixed32_3": uint32(932)},
					PrivateMapStringFixed64_:  map[string]uint64{"PrivateMapStringFixed64_1": uint64(1012), "PrivateMapStringFixed64_2": uint64(1022), "PrivateMapStringFixed64_3": uint64(1032)},
					GrantedMapStringSfixed32_: map[string]int32{"GrantedMapStringSfixed32_1": int32(1112), "GrantedMapStringSfixed32_2": int32(1122), "GrantedMapStringSfixed32_3": int32(1132)},
					PublicMapStringSfixed64_:  map[string]int64{"PublicMapStringSfixed64_1": int64(1212), "PublicMapStringSfixed64_2": int64(1222), "PublicMapStringSfixed64_3": int64(1232)},
					PrivateMapStringBool_:     map[string]bool{"PrivateMapStringBool_1": bool(false), "PrivateMapStringBool_2": bool(true), "PrivateMapStringBool_3": bool(false)},
					GrantedMapStringString_:   map[string]string{"GrantedMapStringString_1": string("1412"), "GrantedMapStringString_2": string("1422"), "GrantedMapStringString_3": string("1432")},
					PublicMapStringBytes_:     map[string][]byte{"PublicMapStringBytes_1": []byte("test"), "PublicMapStringBytes_2": []byte("is"), "PublicMapStringBytes_3": []byte("this")},
					PrivateMapStringMessage_: map[string]*testdatapb.ScalarTypes{
						"PrivateMapStringMessage_1": {PrivateDouble_: 1.12, GrantedFloat_: 2.12, PublicInt32_: 312},
						"PrivateMapStringMessage_2": {PrivateDouble_: 1.22, GrantedFloat_: 2.22, PublicInt32_: 322},
						"PrivateMapStringMessage_3": {PrivateDouble_: 1.32, GrantedFloat_: 2.32, PublicInt32_: 332},
					},
				},
			},
			want: []*compare.Diff{
				{Path: "privateMapStringDouble",
					ValueX: map[string]float64{"PrivateMapStringDouble_1": float64(1.1), "PrivateMapStringDouble_2": float64(1.2), "PrivateMapStringDouble_3": float64(1.3)},
					ValueY: map[string]float64{"PrivateMapStringDouble_1": float64(1.12), "PrivateMapStringDouble_2": float64(1.22), "PrivateMapStringDouble_3": float64(1.32)},
				},
				{Path: "grantedMapStringFloat",
					ValueX: map[string]float32{"GrantedMapStringFloat_1": float32(2.1), "GrantedMapStringFloat_2": float32(2.2), "GrantedMapStringFloat_3": float32(2.3)},
					ValueY: map[string]float32{"GrantedMapStringFloat_1": float32(2.12), "GrantedMapStringFloat_2": float32(2.22), "GrantedMapStringFloat_3": float32(2.32)},
				},
				{Path: "publicMapStringInt32",
					ValueX: map[string]int32{"PublicMapStringInt32_1": int32(31), "PublicMapStringInt32_2": int32(32), "PublicMapStringInt32_3": int32(33)},
					ValueY: map[string]int32{"PublicMapStringInt32_1": int32(312), "PublicMapStringInt32_2": int32(322), "PublicMapStringInt32_3": int32(332)},
				},
				{Path: "privateMapStringInt64",
					ValueX: map[string]int64{"PrivateMapStringInt64_1": int64(41), "PrivateMapStringInt64_2": int64(42), "PrivateMapStringInt64_3": int64(43)},
					ValueY: map[string]int64{"PrivateMapStringInt64_1": int64(412), "PrivateMapStringInt64_2": int64(422), "PrivateMapStringInt64_3": int64(432)},
				},
				{Path: "grantedMapStringUint32",
					ValueX: map[string]uint32{"GrantedMapStringUint32_1": uint32(51), "GrantedMapStringUint32_2": uint32(52), "GrantedMapStringUint32_3": uint32(53)},
					ValueY: map[string]uint32{"GrantedMapStringUint32_1": uint32(512), "GrantedMapStringUint32_2": uint32(522), "GrantedMapStringUint32_3": uint32(532)},
				},
				{Path: "publicMapStringUint64",
					ValueX: map[string]uint64{"PublicMapStringUint64_1": uint64(61), "PublicMapStringUint64_2": uint64(62), "PublicMapStringUint64_3": uint64(63)},
					ValueY: map[string]uint64{"PublicMapStringUint64_1": uint64(612), "PublicMapStringUint64_2": uint64(622), "PublicMapStringUint64_3": uint64(632)},
				},
				{Path: "privateMapStringSint32",
					ValueX: map[string]int32{"PrivateMapStringSint32_1": int32(71), "PrivateMapStringSint32_2": int32(72), "PrivateMapStringSint32_3": int32(73)},
					ValueY: map[string]int32{"PrivateMapStringSint32_1": int32(712), "PrivateMapStringSint32_2": int32(722), "PrivateMapStringSint32_3": int32(732)},
				},
				{Path: "grantedMapStringSint64",
					ValueX: map[string]int64{"GrantedMapStringSint64_1": int64(81), "GrantedMapStringSint64_2": int64(82), "GrantedMapStringSint64_3": int64(83)},
					ValueY: map[string]int64{"GrantedMapStringSint64_1": int64(812), "GrantedMapStringSint64_2": int64(822), "GrantedMapStringSint64_3": int64(832)},
				},
				{Path: "publicMapStringFixed32",
					ValueX: map[string]uint32{"PublicMapStringFixed32_1": uint32(91), "PublicMapStringFixed32_2": uint32(92), "PublicMapStringFixed32_3": uint32(93)},
					ValueY: map[string]uint32{"PublicMapStringFixed32_1": uint32(912), "PublicMapStringFixed32_2": uint32(922), "PublicMapStringFixed32_3": uint32(932)},
				},
				{Path: "privateMapStringFixed64",
					ValueX: map[string]uint64{"PrivateMapStringFixed64_1": uint64(101), "PrivateMapStringFixed64_2": uint64(102), "PrivateMapStringFixed64_3": uint64(103)},
					ValueY: map[string]uint64{"PrivateMapStringFixed64_1": uint64(1012), "PrivateMapStringFixed64_2": uint64(1022), "PrivateMapStringFixed64_3": uint64(1032)},
				},
				{Path: "grantedMapStringSfixed32",
					ValueX: map[string]int32{"GrantedMapStringSfixed32_1": int32(111), "GrantedMapStringSfixed32_2": int32(112), "GrantedMapStringSfixed32_3": int32(113)},
					ValueY: map[string]int32{"GrantedMapStringSfixed32_1": int32(1112), "GrantedMapStringSfixed32_2": int32(1122), "GrantedMapStringSfixed32_3": int32(1132)},
				},
				{Path: "publicMapStringSfixed64",
					ValueX: map[string]int64{"PublicMapStringSfixed64_1": int64(121), "PublicMapStringSfixed64_2": int64(122), "PublicMapStringSfixed64_3": int64(123)},
					ValueY: map[string]int64{"PublicMapStringSfixed64_1": int64(1212), "PublicMapStringSfixed64_2": int64(1222), "PublicMapStringSfixed64_3": int64(1232)},
				},
				{Path: "privateMapStringBool",
					ValueX: map[string]bool{"PrivateMapStringBool_1": bool(true), "PrivateMapStringBool_2": bool(false), "PrivateMapStringBool_3": bool(true)},
					ValueY: map[string]bool{"PrivateMapStringBool_1": bool(false), "PrivateMapStringBool_2": bool(true), "PrivateMapStringBool_3": bool(false)},
				},
				{Path: "grantedMapStringString",
					ValueX: map[string]string{"GrantedMapStringString_1": string("141"), "GrantedMapStringString_2": string("142"), "GrantedMapStringString_3": string("143")},
					ValueY: map[string]string{"GrantedMapStringString_1": string("1412"), "GrantedMapStringString_2": string("1422"), "GrantedMapStringString_3": string("1432")},
				},
				{Path: "publicMapStringBytes",
					ValueX: map[string][]byte{"PublicMapStringBytes_1": []byte("this"), "PublicMapStringBytes_2": []byte("is"), "PublicMapStringBytes_3": []byte("test")},
					ValueY: map[string][]byte{"PublicMapStringBytes_1": []byte("test"), "PublicMapStringBytes_2": []byte("is"), "PublicMapStringBytes_3": []byte("this")},
				},
				{Path: "privateMapStringMessage",
					ValueX: map[string]*testdatapb.ScalarTypes{
						"PrivateMapStringMessage_1": {PrivateDouble_: 1.1, GrantedFloat_: 2.1, PublicInt32_: 31, GrantedString_: "valueX_GrantedString_1"},
						"PrivateMapStringMessage_2": {PrivateDouble_: 1.2, GrantedFloat_: 2.2, PublicInt32_: 32},
						"PrivateMapStringMessage_3": {PrivateDouble_: 1.3, GrantedFloat_: 2.3, PublicInt32_: 33},
					},
					ValueY: map[string]*testdatapb.ScalarTypes{
						"PrivateMapStringMessage_1": {PrivateDouble_: 1.12, GrantedFloat_: 2.12, PublicInt32_: 312},
						"PrivateMapStringMessage_2": {PrivateDouble_: 1.22, GrantedFloat_: 2.22, PublicInt32_: 322},
						"PrivateMapStringMessage_3": {PrivateDouble_: 1.32, GrantedFloat_: 2.32, PublicInt32_: 332},
					},
				},
			},
		},
		{
			name: "one_of_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.OneOfTypes{},
				y:     &testdatapb.OneOfTypes{},
			},
			want: noDiff,
		},
		{
			name: "one_of_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.OneOfTypes{
					TestOneof: &testdatapb.OneOfTypes_PrivateMessage{
						PrivateMessage: &testdatapb.ScalarTypes{PrivateDouble_: 1.1, GrantedFloat_: 2.1, PublicInt32_: 31, GrantedString_: "valueX_GrantedString_1"},
					},
				},
				y: &testdatapb.OneOfTypes{
					TestOneof: &testdatapb.OneOfTypes_GrantedMessage{
						GrantedMessage: &testdatapb.ScalarTypes{PrivateDouble_: 1.12, GrantedFloat_: 2.12, PublicInt32_: 312},
					},
				},
			},
			want: []*compare.Diff{
				{Path: "privateMessage.privateDouble", ValueX: float64(1.1), ValueY: float64(0)},
				{Path: "privateMessage.grantedFloat", ValueX: float32(2.1), ValueY: float32(0)},
				{Path: "privateMessage.publicInt32", ValueX: int32(31), ValueY: int32(0)},
				{Path: "privateMessage.grantedString", ValueX: string("valueX_GrantedString_1"), ValueY: string("")},
				{Path: "grantedMessage.privateDouble", ValueX: float64(0), ValueY: float64(1.12)},
				{Path: "grantedMessage.grantedFloat", ValueX: float32(0), ValueY: float32(2.12)},
				{Path: "grantedMessage.publicInt32", ValueX: int32(0), ValueY: int32(312)},
			},
		},
		{
			name: "nested_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.NestedTypes{},
				y:     &testdatapb.NestedTypes{},
			},
			want: noDiff,
		},
		{
			name: "nested_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.NestedTypes{
					Private:                 string("valueX_Private"),
					Granted:                 string("valueX_Granted"),
					Public:                  string("valueX_Public"),
					PrivateScalarTypes:      &testdatapb.ScalarTypes{GrantedString_: "valueX_PrivateScalarTypes.GrantedString_"},
					GrantedScalarTypes:      &testdatapb.ScalarTypes{GrantedString_: "valueX_GrantedScalarTypes.GrantedString_"},
					PublicScalarTypes:       &testdatapb.ScalarTypes{GrantedString_: "valueX_PublicScalarTypes.GrantedString_"},
					PrivateRepeatedTypes:    &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueX_PrivateRepeatedTypes.GrantedRepeatedString_1", "valueX_PrivateRepeatedTypes.GrantedRepeatedString_2"}},
					GrantedRepeatedTypes:    &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueX_GrantedRepeatedTypes.GrantedRepeatedString_1", "valueX_GrantedRepeatedTypes.GrantedRepeatedString_2"}},
					PublicRepeatedTypes:     &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueX_PublicRepeatedTypes.GrantedRepeatedString_1", "valueX_PublicRepeatedTypes.GrantedRepeatedString_2"}},
					PrivateMapTypes:         &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueX_PrivateMapTypes.PrivateMapStringInt64_1": 1, "valueX_PrivateMapTypes.PrivateMapStringInt64_2": 2}},
					GrantedMapTypes:         &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueX_GrantedMapTypes.PrivateMapStringInt64_1": 1, "valueX_GrantedMapTypes.PrivateMapStringInt64_2": 2}},
					PublicMapTypes:          &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueX_PublicMapTypes.PrivateMapStringInt64_1": 1, "valueX_PublicMapTypes.PrivateMapStringInt64_2": 2}},
					PrivateOneOfTypes:       &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueX_PrivateOneOfTypes.PrivateMessage.GrantedString_"}}},
					GrantedOneOfTypes:       &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueX_GrantedOneOfTypes.PrivateMessage.GrantedString_"}}},
					PublicOneOfTypes:        &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueX_PublicOneOfTypes.PrivateMessage.GrantedString_"}}},
					NestedVisibilityTypesId: string("valueX_NestedVisibilityTypesId"),
					DeletedAt:               int64(1),
					CreatedAt:               int64(2),
					UpdatedAt:               int64(3),
				},
				y: &testdatapb.NestedTypes{
					Private:                 string("valueY_Private"),
					Granted:                 string("valueY_Granted"),
					Public:                  string("valueY_Public"),
					PrivateScalarTypes:      &testdatapb.ScalarTypes{GrantedString_: "valueY_PrivateScalarTypes.GrantedString_"},
					GrantedScalarTypes:      &testdatapb.ScalarTypes{GrantedString_: "valueY_GrantedScalarTypes.GrantedString_"},
					PublicScalarTypes:       &testdatapb.ScalarTypes{GrantedString_: "valueY_PublicScalarTypes.GrantedString_"},
					PrivateRepeatedTypes:    &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueY_PrivateRepeatedTypes.GrantedRepeatedString_1", "valueY_PrivateRepeatedTypes.GrantedRepeatedString_2"}},
					GrantedRepeatedTypes:    &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueY_GrantedRepeatedTypes.GrantedRepeatedString_1", "valueY_GrantedRepeatedTypes.GrantedRepeatedString_2"}},
					PublicRepeatedTypes:     &testdatapb.RepeatedTypes{GrantedRepeatedString_: []string{"valueY_PublicRepeatedTypes.GrantedRepeatedString_1", "valueY_PublicRepeatedTypes.GrantedRepeatedString_2"}},
					PrivateMapTypes:         &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueY_PrivateMapTypes.PrivateMapStringInt64_1": 1, "valueY_PrivateMapTypes.PrivateMapStringInt64_2": 2}},
					GrantedMapTypes:         &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueY_GrantedMapTypes.PrivateMapStringInt64_1": 1, "valueY_GrantedMapTypes.PrivateMapStringInt64_2": 2}},
					PublicMapTypes:          &testdatapb.MapTypes{PrivateMapStringInt64_: map[string]int64{"valueY_PublicMapTypes.PrivateMapStringInt64_1": 1, "valueY_PublicMapTypes.PrivateMapStringInt64_2": 2}},
					PrivateOneOfTypes:       &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueY_PrivateOneOfTypes.PrivateMessage.GrantedString_"}}},
					GrantedOneOfTypes:       &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueY_GrantedOneOfTypes.PrivateMessage.GrantedString_"}}},
					PublicOneOfTypes:        &testdatapb.OneOfTypes{TestOneof: &testdatapb.OneOfTypes_PrivateMessage{PrivateMessage: &testdatapb.ScalarTypes{GrantedString_: "valueY_PublicOneOfTypes.PrivateMessage.GrantedString_"}}},
					NestedVisibilityTypesId: string("valueY_NestedVisibilityTypesId"),
					DeletedAt:               int64(12),
					CreatedAt:               int64(22),
					UpdatedAt:               int64(32),
				},
			},
			want: []*compare.Diff{
				{Path: "private",
					ValueX: string("valueX_Private"),
					ValueY: string("valueY_Private")},
				{Path: "granted",
					ValueX: string("valueX_Granted"),
					ValueY: string("valueY_Granted"),
				},
				{Path: "public",
					ValueX: string("valueX_Public"),
					ValueY: string("valueY_Public"),
				},
				{Path: "privateScalarTypes.grantedString",
					ValueX: string("valueX_PrivateScalarTypes.GrantedString_"),
					ValueY: string("valueY_PrivateScalarTypes.GrantedString_"),
				},
				{Path: "grantedScalarTypes.grantedString",
					ValueX: string("valueX_GrantedScalarTypes.GrantedString_"),
					ValueY: string("valueY_GrantedScalarTypes.GrantedString_"),
				},
				{Path: "publicScalarTypes.grantedString",
					ValueX: string("valueX_PublicScalarTypes.GrantedString_"),
					ValueY: string("valueY_PublicScalarTypes.GrantedString_"),
				},
				{Path: "privateRepeatedTypes.grantedRepeatedString",
					ValueX: []string{"valueX_PrivateRepeatedTypes.GrantedRepeatedString_1", "valueX_PrivateRepeatedTypes.GrantedRepeatedString_2"},
					ValueY: []string{"valueY_PrivateRepeatedTypes.GrantedRepeatedString_1", "valueY_PrivateRepeatedTypes.GrantedRepeatedString_2"},
				},
				{Path: "grantedRepeatedTypes.grantedRepeatedString",
					ValueX: []string{"valueX_GrantedRepeatedTypes.GrantedRepeatedString_1", "valueX_GrantedRepeatedTypes.GrantedRepeatedString_2"},
					ValueY: []string{"valueY_GrantedRepeatedTypes.GrantedRepeatedString_1", "valueY_GrantedRepeatedTypes.GrantedRepeatedString_2"},
				},
				{Path: "publicRepeatedTypes.grantedRepeatedString",
					ValueX: []string{"valueX_PublicRepeatedTypes.GrantedRepeatedString_1", "valueX_PublicRepeatedTypes.GrantedRepeatedString_2"},
					ValueY: []string{"valueY_PublicRepeatedTypes.GrantedRepeatedString_1", "valueY_PublicRepeatedTypes.GrantedRepeatedString_2"},
				},
				{Path: "privateMapTypes.privateMapStringInt64",
					ValueX: map[string]int64{"valueX_PrivateMapTypes.PrivateMapStringInt64_1": 1, "valueX_PrivateMapTypes.PrivateMapStringInt64_2": 2},
					ValueY: map[string]int64{"valueY_PrivateMapTypes.PrivateMapStringInt64_1": 1, "valueY_PrivateMapTypes.PrivateMapStringInt64_2": 2},
				},
				{Path: "grantedMapTypes.privateMapStringInt64",
					ValueX: map[string]int64{"valueX_GrantedMapTypes.PrivateMapStringInt64_1": 1, "valueX_GrantedMapTypes.PrivateMapStringInt64_2": 2},
					ValueY: map[string]int64{"valueY_GrantedMapTypes.PrivateMapStringInt64_1": 1, "valueY_GrantedMapTypes.PrivateMapStringInt64_2": 2},
				},
				{Path: "publicMapTypes.privateMapStringInt64",
					ValueX: map[string]int64{"valueX_PublicMapTypes.PrivateMapStringInt64_1": 1, "valueX_PublicMapTypes.PrivateMapStringInt64_2": 2},
					ValueY: map[string]int64{"valueY_PublicMapTypes.PrivateMapStringInt64_1": 1, "valueY_PublicMapTypes.PrivateMapStringInt64_2": 2},
				},
				{Path: "privateOneOfTypes.privateMessage.grantedString",
					ValueX: string("valueX_PrivateOneOfTypes.PrivateMessage.GrantedString_"),
					ValueY: string("valueY_PrivateOneOfTypes.PrivateMessage.GrantedString_"),
				},
				{Path: "grantedOneOfTypes.privateMessage.grantedString",
					ValueX: string("valueX_GrantedOneOfTypes.PrivateMessage.GrantedString_"),
					ValueY: string("valueY_GrantedOneOfTypes.PrivateMessage.GrantedString_"),
				},
				{Path: "publicOneOfTypes.privateMessage.grantedString",
					ValueX: string("valueX_PublicOneOfTypes.PrivateMessage.GrantedString_"),
					ValueY: string("valueY_PublicOneOfTypes.PrivateMessage.GrantedString_"),
				},
				{Path: "nestedVisibilityTypesId",
					ValueX: string("valueX_NestedVisibilityTypesId"),
					ValueY: string("valueY_NestedVisibilityTypesId"),
				},
				{Path: "deletedAt",
					ValueX: int64(1),
					ValueY: int64(12),
				},
				{Path: "createdAt",
					ValueX: int64(2),
					ValueY: int64(22),
				},
				{Path: "updatedAt",
					ValueX: int64(3),
					ValueY: int64(32),
				},
			},
		},
		{
			name: "nested_inherit_types_zero_value",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x:     &testdatapb.NestedInheritTypes{},
				y:     &testdatapb.NestedInheritTypes{},
			},
			want: noDiff,
		},
		{
			name: "nested_inherit_types_private",
			args: args{
				scope: mtnpb.Scope_PRIVATE,
				x: &testdatapb.NestedInheritTypes{
					PrivateInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueX_PrivateInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PrivateInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PrivateInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PrivateInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PrivateInheritMessage.InheritNestedMessage.InheritString"},
					},
					GrantedInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueX_GrantedInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_GrantedInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_GrantedInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_GrantedInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_GrantedInheritMessage.InheritNestedMessage.InheritString"},
					},
					PublicInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueX_PublicInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PublicInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PublicInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PublicInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueX_PublicInheritMessage.InheritNestedMessage.InheritString"},
					},
				},
				y: &testdatapb.NestedInheritTypes{
					PrivateInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueY_PrivateInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PrivateInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PrivateInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PrivateInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PrivateInheritMessage.InheritNestedMessage.InheritString"},
					},
					GrantedInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueY_GrantedInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_GrantedInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_GrantedInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_GrantedInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_GrantedInheritMessage.InheritNestedMessage.InheritString"},
					},
					PublicInheritMessage: &testdatapb.InheritMessage{
						InheritString:        "valueY_PublicInheritMessage.InheritString",
						PrivateNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PublicInheritMessage.PrivateNestedMessage.InheritString"},
						GrantedNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PublicInheritMessage.GrantedNestedMessage.InheritString"},
						PublicNestedMessage:  &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PublicInheritMessage.PublicNestedMessage.InheritString"},
						InheritNestedMessage: &testdatapb.InheritMessage_NestedMessage{InheritString: "valueY_PublicInheritMessage.InheritNestedMessage.InheritString"},
					},
				},
			},
			want: []*compare.Diff{
				{Path: "privateInheritMessage.inheritString", ValueX: "valueX_PrivateInheritMessage.InheritString", ValueY: "valueY_PrivateInheritMessage.InheritString"},
				{Path: "privateInheritMessage.privateNestedMessage.inheritString", ValueX: "valueX_PrivateInheritMessage.PrivateNestedMessage.InheritString", ValueY: "valueY_PrivateInheritMessage.PrivateNestedMessage.InheritString"},
				{Path: "privateInheritMessage.grantedNestedMessage.inheritString", ValueX: "valueX_PrivateInheritMessage.GrantedNestedMessage.InheritString", ValueY: "valueY_PrivateInheritMessage.GrantedNestedMessage.InheritString"},
				{Path: "privateInheritMessage.publicNestedMessage.inheritString", ValueX: "valueX_PrivateInheritMessage.PublicNestedMessage.InheritString", ValueY: "valueY_PrivateInheritMessage.PublicNestedMessage.InheritString"},
				{Path: "privateInheritMessage.inheritNestedMessage.inheritString", ValueX: "valueX_PrivateInheritMessage.InheritNestedMessage.InheritString", ValueY: "valueY_PrivateInheritMessage.InheritNestedMessage.InheritString"},
				{Path: "grantedInheritMessage.inheritString", ValueX: "valueX_GrantedInheritMessage.InheritString", ValueY: "valueY_GrantedInheritMessage.InheritString"},
				{Path: "grantedInheritMessage.privateNestedMessage.inheritString", ValueX: "valueX_GrantedInheritMessage.PrivateNestedMessage.InheritString", ValueY: "valueY_GrantedInheritMessage.PrivateNestedMessage.InheritString"},
				{Path: "grantedInheritMessage.grantedNestedMessage.inheritString", ValueX: "valueX_GrantedInheritMessage.GrantedNestedMessage.InheritString", ValueY: "valueY_GrantedInheritMessage.GrantedNestedMessage.InheritString"},
				{Path: "grantedInheritMessage.publicNestedMessage.inheritString", ValueX: "valueX_GrantedInheritMessage.PublicNestedMessage.InheritString", ValueY: "valueY_GrantedInheritMessage.PublicNestedMessage.InheritString"},
				{Path: "grantedInheritMessage.inheritNestedMessage.inheritString", ValueX: "valueX_GrantedInheritMessage.InheritNestedMessage.InheritString", ValueY: "valueY_GrantedInheritMessage.InheritNestedMessage.InheritString"},
				{Path: "publicInheritMessage.inheritString", ValueX: "valueX_PublicInheritMessage.InheritString", ValueY: "valueY_PublicInheritMessage.InheritString"},
				{Path: "publicInheritMessage.privateNestedMessage.inheritString", ValueX: "valueX_PublicInheritMessage.PrivateNestedMessage.InheritString", ValueY: "valueY_PublicInheritMessage.PrivateNestedMessage.InheritString"},
				{Path: "publicInheritMessage.grantedNestedMessage.inheritString", ValueX: "valueX_PublicInheritMessage.GrantedNestedMessage.InheritString", ValueY: "valueY_PublicInheritMessage.GrantedNestedMessage.InheritString"},
				{Path: "publicInheritMessage.publicNestedMessage.inheritString", ValueX: "valueX_PublicInheritMessage.PublicNestedMessage.InheritString", ValueY: "valueY_PublicInheritMessage.PublicNestedMessage.InheritString"},
				{Path: "publicInheritMessage.inheritNestedMessage.inheritString", ValueX: "valueX_PublicInheritMessage.InheritNestedMessage.InheritString", ValueY: "valueY_PublicInheritMessage.InheritNestedMessage.InheritString"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compare.Equal(tt.args.scope, tt.args.x, tt.args.y)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			if !assert.Len(t, got, len(tt.want)) {
				return
			}
			for i, v := range got {
				exp := tt.want[i]
				switch v.Path {
				case "privateRepeatedMessage",
					"privateMapStringMessage":
					// 内部フィールドなどは比較したくないため、一度JSON化して比較する
					assert.JSONEqf(t, jsonStr(exp), jsonStr(v), "path=%s", v.Path)
				default:
					assert.Equalf(t, exp, v, "path=%s", v.Path)
				}
			}
		})
	}
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
