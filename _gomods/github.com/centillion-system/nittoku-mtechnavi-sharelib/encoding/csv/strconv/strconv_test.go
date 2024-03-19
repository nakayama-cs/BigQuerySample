package strconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestToString(t *testing.T) {
	type args struct {
		k protoreflect.Kind
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "float64",
			args: args{
				k: protoreflect.DoubleKind,
				v: float64(1.1),
			},
			want: "1.1",
		},
		{
			name: "float32",
			args: args{
				k: protoreflect.FloatKind,
				v: float32(2.1),
			},
			want: "2.1",
		},
		{
			name: "int",
			args: args{
				k: protoreflect.Int64Kind,
				v: int(3),
			},
			want: "3",
		},
		{
			name: "int32",
			args: args{
				k: protoreflect.Int32Kind,
				v: int32(4),
			},
			want: "4",
		},
		{
			name: "int64",
			args: args{
				k: protoreflect.Int64Kind,
				v: int64(5),
			},
			want: "5",
		},
		{
			name: "uint32",
			args: args{
				k: protoreflect.Uint32Kind,
				v: uint32(6),
			},
			want: "6",
		},
		{
			name: "uint64",
			args: args{
				k: protoreflect.Uint64Kind,
				v: uint64(7),
			},
			want: "7",
		},
		{
			name: "bool_true",
			args: args{
				k: protoreflect.BoolKind,
				v: bool(true),
			},
			want: "1",
		},
		{
			name: "bool_false",
			args: args{
				k: protoreflect.BoolKind,
				v: bool(false),
			},
			want: "0",
		},
		{
			name: "string",
			args: args{
				k: protoreflect.StringKind,
				v: string("9"),
			},
			want: "9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.k, tt.args.v)
			if tt.wantErr {
				assert.Errorf(t, err, "ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "ToString() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		k protoreflect.Kind
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		want1   bool
		wantErr bool
	}{
		{
			name: "float64",
			args: args{
				k: protoreflect.DoubleKind,
				v: "1.1",
			},
			want:  float64(1.1),
			want1: true,
		},
		{
			name: "float32",
			args: args{
				k: protoreflect.FloatKind,
				v: "2.1",
			},
			want:  float32(2.1),
			want1: true,
		},
		{
			name: "int32",
			args: args{
				k: protoreflect.Int32Kind,
				v: "4",
			},
			want:  int32(4),
			want1: true,
		},
		{
			name: "int64",
			args: args{
				k: protoreflect.Int64Kind,
				v: "5",
			},
			want:  int64(5),
			want1: true,
		},
		{
			name: "uint32",
			args: args{
				k: protoreflect.Uint32Kind,
				v: "6",
			},
			want:  uint32(6),
			want1: true,
		},
		{
			name: "uint64",
			args: args{
				k: protoreflect.Uint64Kind,
				v: "7",
			},
			want:  uint64(7),
			want1: true,
		},
		{
			name: "bool_true",
			args: args{
				k: protoreflect.BoolKind,
				v: "true",
			},
			wantErr: true,
		},
		{
			name: "bool_false",
			args: args{
				k: protoreflect.BoolKind,
				v: "false",
			},
			wantErr: true,
		},
		{
			name: "bool_0",
			args: args{
				k: protoreflect.BoolKind,
				v: "0",
			},
			want:  bool(false),
			want1: false,
		},
		{
			name: "bool_1",
			args: args{
				k: protoreflect.BoolKind,
				v: "1",
			},
			want:  bool(true),
			want1: true,
		},
		{
			name: "string",
			args: args{
				k: protoreflect.StringKind,
				v: "9",
			},
			want:  string("9"),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := FromString(tt.args.k, tt.args.v)
			if tt.wantErr {
				assert.Errorf(t, err, "FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "FromString() got = %v, want %v", got, tt.want)
			assert.Equalf(t, tt.want1, got1, "FromString() got1 = %v, want1 %v", got1, tt.want1)
		})
	}
}

func TestToCSV(t *testing.T) {
	type args struct {
		k protoreflect.Kind
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "slice_string",
			args: args{
				k: protoreflect.StringKind,
				v: []string{"a1", "a2", "a3", "a4"},
			},
			want: "a1,a2,a3,a4",
		},
		{
			name: "slice_float64",
			args: args{
				k: protoreflect.DoubleKind,
				v: []float64{1.1, 2.1, 3.1, 4.1},
			},
			want: "1.1,2.1,3.1,4.1",
		},
		{
			name: "slice_float32",
			args: args{
				k: protoreflect.FloatKind,
				v: []float64{1.2, 2.2, 3.2, 4.2},
			},
			want: "1.2,2.2,3.2,4.2",
		},
		{
			name: "slice_int64",
			args: args{
				k: protoreflect.Int64Kind,
				v: []int64{11, 21, 31, 41},
			},
			want: "11,21,31,41",
		},
		{
			name: "slice_int32",
			args: args{
				k: protoreflect.Int32Kind,
				v: []int32{12, 22, 32, 42},
			},
			want: "12,22,32,42",
		},
		{
			name: "slice_bool_1",
			args: args{
				k: protoreflect.BoolKind,
				v: []bool{true, false, true, false, false},
			},
			want: "1,0,1,0,0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCSV(tt.args.k, tt.args.v)
			if tt.wantErr {
				assert.Errorf(t, err, "ToCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "ToCSV() = %v, want %v", got, tt.want)
		})
	}
}

func TestFromCSV(t *testing.T) {
	type args struct {
		k protoreflect.Kind
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		want1   bool
		wantErr bool
	}{
		{
			name: "slice_string",
			args: args{
				k: protoreflect.StringKind,
				v: "a1,a2,a3,a4",
			},
			want:  []string{"a1", "a2", "a3", "a4"},
			want1: true,
		},
		{
			name: "slice_float64",
			args: args{
				k: protoreflect.DoubleKind,
				v: "1.1,2.1,3.1,4.1",
			},
			want:  []float64{1.1, 2.1, 3.1, 4.1},
			want1: true,
		},
		{
			name: "slice_float32",
			args: args{
				k: protoreflect.FloatKind,
				v: "1.2,2.2,3.2,4.2",
			},
			want:  []float32{1.2, 2.2, 3.2, 4.2},
			want1: true,
		},
		{
			name: "slice_int64",
			args: args{
				k: protoreflect.Int64Kind,
				v: "11,21,31,41",
			},
			want:  []int64{11, 21, 31, 41},
			want1: true,
		},
		{
			name: "slice_int32",
			args: args{
				k: protoreflect.Int32Kind,
				v: "12,22,32,42",
			},
			want:  []int32{12, 22, 32, 42},
			want1: true,
		},
		{
			name: "slice_bool",
			args: args{
				k: protoreflect.BoolKind,
				v: "1,0,1,0,0",
			},
			want:  []bool{true, false, true, false, false},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := FromCSV(tt.args.k, tt.args.v)

			if tt.wantErr {
				assert.Errorf(t, err, "FromCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "FromCSV() got = %v, want %v", got, tt.want)
			assert.Equalf(t, tt.want1, got1, "FromCSV() got1 = %v, want1 %v", got1, tt.want1)
		})
	}
}

func Test_makeSlice(t *testing.T) {
	type args struct {
		k        protoreflect.Kind
		length   int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "slice_string",
			args: args{
				k:        protoreflect.StringKind,
				length:   0,
				capacity: 3,
			},
			want: []string{},
		},
		{
			name: "slice_float64",
			args: args{
				k:        protoreflect.DoubleKind,
				length:   0,
				capacity: 3,
			},
			want: []float64{},
		},
		{
			name: "slice_float32",
			args: args{
				k:        protoreflect.FloatKind,
				length:   0,
				capacity: 3,
			},
			want: []float32{},
		},
		{
			name: "slice_int64",
			args: args{
				k:        protoreflect.Int64Kind,
				length:   0,
				capacity: 3,
			},
			want: []int64{},
		},
		{
			name: "slice_int32",
			args: args{
				k:        protoreflect.Int32Kind,
				length:   0,
				capacity: 3,
			},
			want: []int32{},
		},
		{
			name: "slice_bool",
			args: args{
				k:        protoreflect.BoolKind,
				length:   0,
				capacity: 3,
			},
			want: []bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeSlice(tt.args.k, tt.args.length, tt.args.capacity)
			assert.Equalf(t, tt.want, got, "makeSlice() = %v, want %v", got, tt.want)
		})
	}
}
