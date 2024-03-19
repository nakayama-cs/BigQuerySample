package ioengine_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mtechnavi/sharelib/constants"
	"mtechnavi/sharelib/encoding/csv"
	"mtechnavi/sharelib/ioengine"
	testdatapb "mtechnavi/sharelib/ioengine/testdata/protobuf"
	fileformatpb "mtechnavi/sharelib/ioengine/testdata/protobuf/fileformat"
	sharelibpb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/validator"
	"strings"
	"time"

	"mtechnavi/sharelib/to"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestExportCSVToSignedURL(t *testing.T) {
	ctx := context.Background()

	tz, err := time.LoadLocation(constants.TimeZoneName_AsiaTokyo)
	if err != nil {
		panic(err)
	}

	type args struct {
		records     []proto.Message
		format      *sharelibpb.FileFormat
		transformer func(proto.Message) csv.Marshaler
	}
	type expect struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		wantErr bool
	}{
		{
			name:    "export_success",
			wantErr: false,
			args: args{
				format: to.ToFileFormat(&fileformatpb.TestRecordFormatDefault{}),
				transformer: func(m proto.Message) csv.Marshaler {
					return toTestRecordFormatDefault(m.(*testdatapb.TestRecord))
				},
				records: []proto.Message{
					&testdatapb.TestRecord{
						Double_:         1.1,
						Float_:          2.1,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
			},
			expect: expect{
				csv: strings.Join([]string{
					`"TEST_RECORD_ID","DOUBLE","FLOAT","INT32","INT64","UINT32","UINT64","SINT32","SINT64","FIXED32","FIXED64","SFIXED32","SFIXED64","BOOL","STRING","","REPEATED_STRING","REPEATED_INT64","UPDATED_AT"`,
					`"","1.1","2.1","3","4","5","6","7","8","9","10","11","12","0","14string","","a,b,c,d","151,152,153","0"`,
					``,
				}, "\r\n"),
			},
		},
		{
			name:    "export_with_error_content",
			wantErr: false,
			args: args{
				format: to.ToFileFormat(&fileformatpb.TestRecordWithError{}),
				transformer: func(m proto.Message) csv.Marshaler {
					return (m.(*fileformatpb.TestRecordWithError))
				},
				records: []proto.Message{
					&fileformatpb.TestRecordWithError{
						Base: &fileformatpb.TestRecordFormatDefault{
							Double_:         1.1,
							Float_:          2.1,
							Int32_:          3,
							Int64_:          4,
							Uint32_:         5,
							Uint64_:         6,
							Sint32_:         7,
							Sint64_:         8,
							Fixed32_:        uint32(9),
							Fixed64_:        uint64(10),
							Sfixed32_:       11,
							Sfixed64_:       12,
							Bool_:           false,
							String_:         "14string",
							RepeatedString_: []string{"a", "b", "c", "d"},
							RepeatedInt64_:  []int64{151, 152, 153},
						},
						ErrorContent: "this,is,error",
					},
				},
			},
			expect: expect{
				csv: strings.Join([]string{
					`"TEST_RECORD_ID","DOUBLE","FLOAT","INT32","INT64","UINT32","UINT64","SINT32","SINT64","FIXED32","FIXED64","SFIXED32","SFIXED64","BOOL","STRING","","REPEATED_STRING","REPEATED_INT64","UPDATED_AT","ERROR_CONTENT"`,
					`"","1.1","2.1","3","4","5","6","7","8","9","10","11","12","0","14string","","a,b,c,d","151,152,153","0","this,is,error"`,
					``,
				}, "\r\n"),
			},
		},
		{
			name:    "export_datetime",
			wantErr: false,
			args: args{
				format: to.ToFileFormat(&fileformatpb.DatetimeRecordFormatDefault{}),
				transformer: func(m proto.Message) csv.Marshaler {
					return toDatetimeRecordFormatDefault(m.(*testdatapb.DatetimeRecord))
				},
				records: []proto.Message{
					&testdatapb.DatetimeRecord{
						Datetime1_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 11, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime3_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 1, 5, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime5_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 11, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime7_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 1, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime9_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 1, 5, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
					},
					&testdatapb.DatetimeRecord{
						Datetime1_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 12, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime3_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 2, 6, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime5_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 12, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime7_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 2, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
						Datetime9_: &sharelibpb.Datetime{
							TimezoneName: constants.TimeZoneName_AsiaTokyo,
							Timestamp:    time.Date(2022, 2, 6, 0, 0, 0, 000000000, tz).UnixMicro(),
							Accuracy:     sharelibpb.Datetime_DAY,
						},
					},
				},
			},
			expect: expect{
				csv: strings.Join([]string{
					`"DATETIME1_","DATETIME3_","DATETIME5_","DATETIME7_","DATETIME9_"`,
					`"2022/11/25","2022/01/05","2022/11/25","2022/01/25","2022/01/05"`,
					`"2022/12/26","2022/02/06","2022/12/26","2022/02/26","2022/02/06"`,
					``,
				}, "\r\n"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
					return
				}
				assert.Equal(t, strings.TrimSpace(tt.expect.csv), strings.TrimSpace(string(b)))
				fmt.Fprintln(w, "ok: this is test")
			}))
			defer ts.Close()
			err := ioengine.ExportCSVToSignedURL(ctx, tt.args.format, ts.URL, tt.args.records, tt.args.transformer)
			if tt.wantErr {
				assert.NotNilf(t, err, "ExportCSVToSignedURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestImportCSVFromSignedURL(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer ts.Close()

	format := to.ToFileFormat(&fileformatpb.TestRecordFormatDefault{})
	transformer := func(m csv.Unmarshaler, base proto.Message) (proto.Message, error) {
		toTestRecord(m.(*fileformatpb.TestRecordFormatDefault), base.(*testdatapb.TestRecord))
		return base, nil
	}

	type args struct {
		signedURL string
		bases     interface{}
		rules     []validator.Rule
	}
	type errorDetail struct {
		message string
		column  int
		line    int
	}
	type expect struct {
		errorContents []string
		message       proto.Message
		errs          []*errorDetail
	}
	tests := []struct {
		name    string
		args    args
		expects []expect
		wantErr bool
	}{
		{
			name: "import_success",
			args: args{
				signedURL: ts.URL + "/success.csv",
				bases:     []*testdatapb.TestRecord{},
				rules:     []validator.Rule{},
			},
			wantErr: false,
			expects: []expect{
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:    "test_record_id_01",
						Double_:         1.1,
						Float_:          2.1,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
						UpdatedAt:       16,
					},
				},
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:    "",
						Double_:         1.2,
						Float_:          2.2,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:    "test_record_id_03",
						Double_:         1.3,
						Float_:          2.3,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
			},
		},
		{
			name: "import_with_base_records",
			args: args{
				signedURL: ts.URL + "/success.csv",
				bases: []*testdatapb.TestRecord{
					{
						TestRecordId:        "test_record_id_01",
						Double_:             1.1234,
						DoubleBase_:         1.2,
						Float_:              2.1234,
						FloatBase_:          2.2,
						Int32_:              31234,
						Int32Base_:          32,
						Int64_:              41234,
						Int64Base_:          42,
						Uint32_:             51234,
						Uint32Base_:         52,
						Uint64_:             61234,
						Uint64Base_:         62,
						Sint32_:             71234,
						Sint32Base_:         72,
						Sint64_:             81234,
						Sint64Base_:         82,
						Fixed32_:            uint32(91234),
						Fixed32Base_:        uint32(92),
						Fixed64_:            uint64(101234),
						Fixed64Base_:        uint64(102),
						Sfixed32_:           111234,
						Sfixed32Base_:       112,
						Sfixed64_:           121234,
						Sfixed64Base_:       122,
						Bool_:               true,
						BoolBase_:           false,
						String_:             "14string_1234",
						StringBase_:         "14string_2",
						RepeatedString_:     []string{"a1234", "b1234", "c1234", "d1234"},
						RepeatedStringBase_: []string{"a2", "b2", "c2", "d2"},
						RepeatedInt64_:      []int64{1511234, 1521234, 1531234},
						RepeatedInt64Base_:  []int64{154, 155, 156},
						UpdatedAt:           int64(1234),
					},
					{
						TestRecordId:    "test_record_id_03",
						Double_:         0,
						Float_:          0,
						Int32_:          int32(0),
						Int64_:          int64(0),
						Uint32_:         0,
						Uint64_:         0,
						Sint32_:         0,
						Sint64_:         0,
						Fixed32_:        uint32(0),
						Fixed64_:        uint64(0),
						Sfixed32_:       0,
						Sfixed64_:       0,
						Bool_:           false,
						String_:         "",
						RepeatedString_: []string{},
						RepeatedInt64_:  []int64{},
						UpdatedAt:       int64(0),
					},
				},
				rules: []validator.Rule{},
			},
			wantErr: false,
			expects: []expect{
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:        "test_record_id_01",
						Double_:             1.1,
						DoubleBase_:         1.2,
						Float_:              2.1,
						FloatBase_:          2.2,
						Int32_:              3,
						Int32Base_:          32,
						Int64_:              4,
						Int64Base_:          42,
						Uint32_:             5,
						Uint32Base_:         52,
						Uint64_:             6,
						Uint64Base_:         62,
						Sint32_:             7,
						Sint32Base_:         72,
						Sint64_:             8,
						Sint64Base_:         82,
						Fixed32_:            uint32(9),
						Fixed32Base_:        uint32(92),
						Fixed64_:            uint64(10),
						Fixed64Base_:        uint64(102),
						Sfixed32_:           11,
						Sfixed32Base_:       112,
						Sfixed64_:           12,
						Sfixed64Base_:       122,
						Bool_:               false,
						BoolBase_:           false,
						String_:             "14string",
						StringBase_:         "14string_2",
						RepeatedString_:     []string{"a", "b", "c", "d"},
						RepeatedStringBase_: []string{"a2", "b2", "c2", "d2"},
						RepeatedInt64_:      []int64{151, 152, 153},
						RepeatedInt64Base_:  []int64{154, 155, 156},
						UpdatedAt:           16,
					},
				},
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:    "",
						Double_:         1.2,
						Float_:          2.2,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
				{
					errs:          nil,
					errorContents: []string(nil),
					message: &testdatapb.TestRecord{
						TestRecordId:    "test_record_id_03",
						Double_:         1.3,
						Float_:          2.3,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
			},
		},
		{
			name: "import_with_csv_error",
			args: args{
				signedURL: ts.URL + "/error.csv",
				bases:     []*testdatapb.TestRecord{},
				rules:     []validator.Rule{},
			},
			wantErr: false,
			expects: []expect{
				{
					errorContents: []string{
						"E1000002",
					},
					errs: []*errorDetail{
						{
							message: strings.Join([]string{
								"E0000004",
							}, ","),
							column: 18,
							line:   3,
						},
					},
					message: &testdatapb.TestRecord{
						TestRecordId:    "test_record_id_01",
						Double_:         1.1,
						Float_:          2.1,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           true,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  nil,
					},
				},
				{
					errorContents: []string{
						"E1000002",
						"E1000002",
						"E1000002",
					},
					errs: []*errorDetail{
						{
							message: strings.Join([]string{
								"E0000004",
							}, ","),
							column: 11,
							line:   4,
						},
						{
							message: strings.Join([]string{
								"E0000004",
							}, ","),
							column: 13,
							line:   5,
						},
						{
							message: strings.Join([]string{
								"E0000017",
							}, ","),
							column: 14,
							line:   6,
						},
					},
					message: &testdatapb.TestRecord{
						TestRecordId:    "test_record_id_02",
						Double_:         1.2,
						Float_:          2.2,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(0),
						Sfixed32_:       11,
						Sfixed64_:       0,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
			},
		},
		{
			name: "import_with_validate",
			args: args{
				signedURL: ts.URL + "/validate.csv",
				bases:     []*testdatapb.TestRecord{},
				rules: []validator.Rule{
					validator.NotEmptyf("string_", sharelibpb.ErrorLevel_ERROR, "%sが未入力", "STRING"),
					validator.NotEmptyf("repeated_int64_", sharelibpb.ErrorLevel_ERROR, "%sが未入力", "REPEATED_INT64"),
					validator.NotEmptyf("repeated_string_", sharelibpb.ErrorLevel_ERROR, "%sが未入力", "REPEATED_STRING"),
				},
			},
			wantErr: false,
			expects: []expect{
				{
					errorContents: []string{
						"E1000001",
						"E1000001",
					},
					errs: []*errorDetail{
						{
							message: strings.Join([]string{
								"E1000001",
							}, ","),
							column: 15,
							line:   3,
						},
						{
							message: strings.Join([]string{
								"E1000001",
							}, ","),
							column: 17,
							line:   3,
						},
					},
					message: &testdatapb.TestRecord{
						TestRecordId:    "",
						Double_:         1.1,
						Float_:          2.1,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           true,
						String_:         "",
						RepeatedString_: nil,
						RepeatedInt64_:  []int64{151, 152, 153},
					},
				},
				{
					errorContents: []string{
						"E1000001",
					},
					errs: []*errorDetail{
						{
							message: strings.Join([]string{
								"E1000001",
							}, ","),
							column: 18,
							line:   4,
						},
					},
					message: &testdatapb.TestRecord{
						TestRecordId:    "",
						Double_:         1.2,
						Float_:          2.2,
						Int32_:          3,
						Int64_:          4,
						Uint32_:         5,
						Uint64_:         6,
						Sint32_:         7,
						Sint64_:         8,
						Fixed32_:        uint32(9),
						Fixed64_:        uint64(10),
						Sfixed32_:       11,
						Sfixed64_:       12,
						Bool_:           false,
						String_:         "14string",
						RepeatedString_: []string{"a", "b", "c", "d"},
						RepeatedInt64_:  nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts []ioengine.ImportOption
			if len(tt.args.rules) > 0 {
				rs := validator.NewRuleSet()
				rs.Add(tt.args.rules...)
				opts = append(opts, ioengine.WithRuleSet(rs))
			}
			var formatDefaults []*fileformatpb.TestRecordFormatDefault
			results, err := ioengine.ImportCSVFromSignedURL(ctx, format,
				tt.args.signedURL,
				formatDefaults,
				tt.args.bases,
				transformer,
				opts...)
			if tt.wantErr {
				assert.NotNil(t, err, "ImportCSVFromSignedURL() err=%v, want notNil", err)
				return
			} else {
				assert.Nil(t, err, "ImportCSVFromSignedURL() err=%v, want nil", err)
			}
			for idx, res := range results {
				fmt.Printf("test idx %d\n", idx)
				exp := tt.expects[idx]
				fmt.Printf("test expects %v\n", exp.message)

				if exp.errs == nil {
					assert.Nilf(t, res.Err, "csv.Error[%d]: got=%v, want nil", idx, res.Err)
					continue
				} else {
					assert.NotNil(t, res.Err)
				}

				fmt.Println(tt.args.signedURL)
				fmt.Println(res.AllErrors())
				if assert.Lenf(t, res.AllErrors(), len(exp.errs), "csv.Error[%d]: len=%v, want %d", idx, len(res.AllErrors()), len(exp.errs)) {
					for idx, e := range res.AllErrors() {
						expErr := exp.errs[idx]
						switch e := e.(type) {
						case *csv.Error:
							assert.Equalf(t, expErr.column, e.Column, "csv.Error[%d]: column=%v, want %d", idx, e.Column, expErr.column)
							assert.Equalf(t, expErr.line, e.Line, "csv.Error[%d]: line=%v, want %d", idx, e.Line, expErr.line)
						default:
							continue
						}
					}
				}
			}
		})
	}
}

func TestImportDatetimeCSVFromSignedURL(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer ts.Close()

	tz, err := time.LoadLocation(constants.TimeZoneName_AsiaTokyo)
	if err != nil {
		panic(err)
	}

	//テストケース
	t.Run("Validation", func(t *testing.T) {

		format := to.ToFileFormat(&fileformatpb.DatetimeRecordFormatDefault{})
		transformer := func(m csv.Unmarshaler, base proto.Message) (proto.Message, error) {
			toDatetimeRecord(m.(*fileformatpb.DatetimeRecordFormatDefault), base.(*testdatapb.DatetimeRecord))
			return base, nil
		}

		type args struct {
			signedURL string
			bases     interface{}
			rules     []validator.Rule
		}
		type errorDetail struct {
			message string
			column  int
			line    int
		}
		type expect struct {
			errorContents []string
			message       proto.Message
			errs          []*errorDetail
		}
		tests := []struct {
			name    string
			args    args
			expects []expect
			wantErr bool
		}{
			{
				name: "import_datetime",
				args: args{
					signedURL: ts.URL + "/datetime.csv",
					bases:     []*testdatapb.DatetimeRecord{},
					rules:     []validator.Rule{},
				},
				wantErr: false,
				expects: []expect{
					{
						// errorContents: []string{},
						// errs:          []*errorDetail{},
						message: &testdatapb.DatetimeRecord{
							Datetime1_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 11, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime3_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 1, 5, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime5_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 11, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime7_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 1, 25, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime9_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 1, 5, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
						},
					},
					{
						// errorContents: []string{},
						// errs: []*errorDetail{},
						message: &testdatapb.DatetimeRecord{
							Datetime1_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 12, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime3_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 2, 6, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime5_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 12, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime7_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 2, 26, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
							Datetime9_: &sharelibpb.Datetime{
								TimezoneName: constants.TimeZoneName_AsiaTokyo,
								Timestamp:    time.Date(2022, 2, 6, 0, 0, 0, 000000000, tz).UnixMicro(),
								Accuracy:     sharelibpb.Datetime_DAY,
							},
						},
					},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var opts []ioengine.ImportOption
				if len(tt.args.rules) > 0 {
					rs := validator.NewRuleSet()
					rs.Add(tt.args.rules...)
					opts = append(opts, ioengine.WithRuleSet(rs))
				}
				var formatDefaults []*fileformatpb.DatetimeRecordFormatDefault
				results, err := ioengine.ImportCSVFromSignedURL(ctx, format,
					tt.args.signedURL,
					formatDefaults,
					tt.args.bases,
					transformer,
					opts...)
				if tt.wantErr {
					assert.NotNil(t, err, "TestImportDatetimeCSVFromSignedURL() err=%v, want notNil", err)
					return
				} else {
					assert.Nil(t, err, "TestImportDatetimeCSVFromSignedURL() err=%v, want nil", err)
				}
				for idx, res := range results {
					exp := tt.expects[idx]

					assert.Equalf(t, exp.errorContents, res.ErrorContents(constants.LanguageName_JA, &sharelibpb.FileFormat{}), "csv.Error[%d]: errorContents=%v, want %v", idx, res.ErrorContents(constants.LanguageName_JA, &sharelibpb.FileFormat{}), exp.errorContents)
					assert.JSONEqf(t, jsonStr(exp.message), jsonStr(res.Message), "csv.Error[%d]: message= %v, want %v", idx, jsonStr(res.Message), jsonStr(exp.message))

					if exp.errs == nil {
						assert.Nilf(t, res.Err, "csv.Error[%d]: got=%v, want nil", idx, res.Err)
						continue
					} else {
						assert.NotNil(t, res.Err)
					}

					if assert.Lenf(t, res.AllErrors(), len(exp.errs), "csv.Error[%d]: len=%v, want %d", idx, len(res.AllErrors()), len(exp.errs)) {
						for idx, e := range res.AllErrors() {
							expErr := exp.errs[idx]
							assert.Equalf(t, expErr.message, e.Error(), "csv.Error[%d]: message=%v, want %d", idx, e.Error(), expErr.message)
							switch e := e.(type) {
							case *csv.Error:
								assert.Equalf(t, expErr.column, e.Column, "csv.Error[%d]: column=%v, want %d", idx, e.Column, expErr.column)
								assert.Equalf(t, expErr.line, e.Line, "csv.Error[%d]: line=%v, want %d", idx, e.Line, expErr.line)
							default:
								continue
							}
						}
					}
				}
			})
		}
	})

	t.Run("Column Error 1", func(t *testing.T) {
		format := to.ToFileFormat(&fileformatpb.DatetimeRecordFormatDefault{})
		transformer := func(m csv.Unmarshaler, base proto.Message) (proto.Message, error) {
			toDatetimeRecord(m.(*fileformatpb.DatetimeRecordFormatDefault), base.(*testdatapb.DatetimeRecord))
			return base, nil
		}

		var opts []ioengine.ImportOption
		var formatDefaults []*fileformatpb.DatetimeRecordFormatDefault
		results, err := ioengine.ImportCSVFromSignedURL(ctx, format,
			ts.URL+"/column-error1.csv",
			formatDefaults,
			[]*testdatapb.DatetimeRecord{},
			transformer,
			opts...)

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Equal(t, ioengine.ErrIoEngineColumnError, err)
	})

	t.Run("Column Error 2", func(t *testing.T) {
		format := to.ToFileFormat(&fileformatpb.DatetimeRecordFormatDefault{})
		transformer := func(m csv.Unmarshaler, base proto.Message) (proto.Message, error) {
			toDatetimeRecord(m.(*fileformatpb.DatetimeRecordFormatDefault), base.(*testdatapb.DatetimeRecord))
			return base, nil
		}

		var opts []ioengine.ImportOption
		var formatDefaults []*fileformatpb.DatetimeRecordFormatDefault
		results, err := ioengine.ImportCSVFromSignedURL(ctx, format,
			ts.URL+"/column-error2.csv",
			formatDefaults,
			[]*testdatapb.DatetimeRecord{},
			transformer,
			opts...)

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Equal(t, ioengine.ErrIoEngineColumnError, err)
	})

}

func toTestRecord(src *fileformatpb.TestRecordFormatDefault, out *testdatapb.TestRecord) {
	out.TestRecordId = src.TestRecordId
	out.Double_ = src.Double_
	out.Float_ = src.Float_
	out.Int32_ = src.Int32_
	out.Int64_ = src.Int64_
	out.Uint32_ = src.Uint32_
	out.Uint64_ = src.Uint64_
	out.Sint32_ = src.Sint32_
	out.Sint64_ = src.Sint64_
	out.Fixed32_ = src.Fixed32_
	out.Fixed64_ = src.Fixed64_
	out.Sfixed32_ = src.Sfixed32_
	out.Sfixed64_ = src.Sfixed64_
	out.Bool_ = src.Bool_
	out.String_ = src.String_
	// out.RootEnum    = src.Double_
	out.RepeatedString_ = src.RepeatedString_
	out.RepeatedInt64_ = src.RepeatedInt64_
	out.UpdatedAt = src.UpdatedAt
}

func toTestRecordFormatDefault(src *testdatapb.TestRecord) *fileformatpb.TestRecordFormatDefault {
	var out fileformatpb.TestRecordFormatDefault
	out.TestRecordId = src.TestRecordId
	out.Double_ = src.Double_
	out.Float_ = src.Float_
	out.Int32_ = src.Int32_
	out.Int64_ = src.Int64_
	out.Uint32_ = src.Uint32_
	out.Uint64_ = src.Uint64_
	out.Sint32_ = src.Sint32_
	out.Sint64_ = src.Sint64_
	out.Fixed32_ = src.Fixed32_
	out.Fixed64_ = src.Fixed64_
	out.Sfixed32_ = src.Sfixed32_
	out.Sfixed64_ = src.Sfixed64_
	out.Bool_ = src.Bool_
	out.String_ = src.String_
	// out.RootEnum = src.RootEnum
	out.RepeatedString_ = src.RepeatedString_
	out.RepeatedInt64_ = src.RepeatedInt64_
	out.UpdatedAt = src.UpdatedAt
	return &out
}

func toDatetimeRecord(src *fileformatpb.DatetimeRecordFormatDefault, out *testdatapb.DatetimeRecord) {
	out.Datetime1_ = src.Datetime1_
	out.Datetime3_ = src.Datetime3_
	out.Datetime5_ = src.Datetime5_
	out.Datetime7_ = src.Datetime7_
	out.Datetime9_ = src.Datetime9_
}

func toDatetimeRecordFormatDefault(src *testdatapb.DatetimeRecord) *fileformatpb.DatetimeRecordFormatDefault {
	var out fileformatpb.DatetimeRecordFormatDefault
	out.Datetime1_ = src.Datetime1_
	out.Datetime3_ = src.Datetime3_
	out.Datetime5_ = src.Datetime5_
	out.Datetime7_ = src.Datetime7_
	out.Datetime9_ = src.Datetime9_
	return &out
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
