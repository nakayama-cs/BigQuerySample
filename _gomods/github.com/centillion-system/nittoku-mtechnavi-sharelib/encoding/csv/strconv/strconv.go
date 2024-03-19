// 本パッケージは、通常のユーザコードで利用しないでください。
// protoc-gen-csvで生成されたコード内で利用されることを想定しています。
package strconv

import (
	"bytes"
	gocsv "encoding/csv"
	"fmt"
	"mtechnavi/sharelib/constants"
	"mtechnavi/sharelib/encoding/csv"
	sharelibpb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/to"
	"reflect"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// Do not use this.
func ToString(k protoreflect.Kind, v any) (string, error) {
	switch k {
	case protoreflect.BoolKind:
		return ToBoolString(v), nil
	default:
		return fmt.Sprint(v), nil
	}
}

// Do not use this.
func ToBoolString(v any) string {
	switch v {
	case true:
		return "1"
	case false:
		return "0"
	default:
		panic(fmt.Sprintf("unsupported bool value: %v", v))
	}
}

// Do not use this.
func FromString(k protoreflect.Kind, s string) (any, bool, error) {
	fn, ok := converters[k]
	if !ok {
		panic(fmt.Sprintf("unsupported type: %v", k))
	}
	v, err := fn(s)
	if err != nil {
		return v, false, err
	}
	rv := reflect.ValueOf(v)
	ok = rv.IsValid() && !rv.IsZero()
	return v, ok, nil
}

type TimeZone string

const (
	asiaTokyo TimeZone = constants.TimeZoneName_AsiaTokyo
)

func (t TimeZone) toLocation() *time.Location {
	tz, err := time.LoadLocation(string(t))
	if err != nil {
		// TODO: 現状は固定値のためにpanicしないが、外部から指定されるようになった場合はpanicせずに
		// 上位関数にエラーを通知する必要がある。
		panic(fmt.Sprintf("failed to load timezone tz=%v: %v", asiaTokyo, err))
	}
	return tz
}

// TODO: タイムゾーンは外部から指定される想定のため、仕様確定までは起動時にデフォルト("Asia/Tokyo")情報を読み込む
var tz = asiaTokyo.toLocation()

// Do not use this.
func ToDatetime(k protoreflect.Kind, v any) (string, error) {
	dt, ok := v.(*sharelibpb.Datetime)
	if !ok {
		return "", fmt.Errorf("unsupported type: %v", k)
	}
	return toDateTime(dt)
}

func toDateTime(dt *sharelibpb.Datetime) (string, error) {
	if dt.Timestamp == 0 {
		return "", nil
	}

	t := time.UnixMicro(dt.Timestamp)
	tz := TimeZone(dt.TimezoneName).toLocation()

	t = t.In(tz)
	switch dt.Accuracy {
	case sharelibpb.Datetime_MICROSECOND:
		t = t.Truncate(time.Microsecond)
	case sharelibpb.Datetime_MILLISECOND:
		t = t.Truncate(time.Millisecond)
	case sharelibpb.Datetime_SECOND:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, tz)
	case sharelibpb.Datetime_MINUTE:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, tz)
	case sharelibpb.Datetime_HOUR:
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, tz)
	case sharelibpb.Datetime_DAY:
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz)
	case sharelibpb.Datetime_MONTH:
		t = time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, tz)
	case sharelibpb.Datetime_YEAR:
		t = time.Date(t.Year(), 0, 0, 0, 0, 0, 0, tz)
	default:
		return "", fmt.Errorf("unsupported accuracy: %v", dt.Accuracy)
	}
	return t.Format("2006/01/02"), nil
}

// Do not use this.
func FromDatetime(k protoreflect.Kind, s string) (any, bool, error) {
	// 空白は許容し、nilにする
	if s == "" {
		return nil, true, nil
	}

	t, ok := convertToTime(s)
	if !ok {
		return nil, false, fmt.Errorf("unsupported time format: %s", s)
	}
	var dt sharelibpb.Datetime
	dt.Accuracy = sharelibpb.Datetime_DAY
	dt.TimezoneName = string(asiaTokyo)
	dt.Timestamp = t.UnixMicro()
	return &dt, ok, nil
}

func ToDate(k protoreflect.Kind, v any) (string, error) {
	on, ok := v.(*sharelibpb.Date)
	if !ok {
		return "", fmt.Errorf("unsupported type: %v", k)
	}
	return toDate(on)
}

func toDate(on *sharelibpb.Date) (string, error) {
	if on == nil ||
		on.Year == 0 ||
		on.Month == 0 ||
		on.Day == 0 {
		return "", nil
	}
	return fmt.Sprintf("%04d/%02d/%02d", on.Year, on.Month, on.Day), nil
}

func FromDate(k protoreflect.Kind, s string) (any, bool, error) {
	// 空白は許容し、nilにする
	if s == "" {
		return nil, true, nil
	}

	t, ok := convertToTime(s)
	if !ok {
		return nil, false, fmt.Errorf("unsupported time format: %s", s)
	}
	var on sharelibpb.Date
	on.Year = int32(t.Year())
	on.Month = int32(t.Month())
	on.Day = int32(t.Day())
	return &on, ok, nil
}

func ToAmount(k protoreflect.Kind, v any) (string, error) {
	amount, ok := v.(*sharelibpb.Amount)
	if !ok {
		return "", fmt.Errorf("unsupported type: %v", k)
	}
	return toAmount(amount)
}

func toAmount(amount *sharelibpb.Amount) (string, error) {
	if amount == nil {
		return "", nil
	}
	v, err := to.ToFloat64(amount.IntegralAmount, amount.FractionalAmount)
	if err != nil {
		return "", fmt.Errorf("unsupported amount: %v", amount)
	}
	return strconv.FormatFloat(v, 'f', -1, 64), nil
}

func FromAmount(k protoreflect.Kind, s string) (any, bool, error) {
	// 空白は許容し、nilにする
	if s == "" {
		return nil, true, nil
	}

	amount, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, false, fmt.Errorf("unsupported amount format: %s", s)
	}
	i, f, err := to.ToFloat64Parts(amount)
	if err != nil {
		return nil, false, fmt.Errorf("unsupported amount: %v", amount)
	}

	out := &sharelibpb.Amount{
		IntegralAmount:   i,
		FractionalAmount: f,
	}
	return out, true, nil
}

// Do not use this.
func ToCSV(k protoreflect.Kind, l any) (string, error) {
	rv := reflect.ValueOf(l)
	if !rv.IsValid() {
		return "", nil
	}
	record := make([]string, rv.Len())
	for i := range record {
		if v, err := ToString(k, rv.Index(i).Interface()); err != nil {
			return "", err
		} else {
			record[i] = v
		}
	}
	var buffer bytes.Buffer
	writer := gocsv.NewWriter(&buffer)
	writer.Comma = ','
	if err := writer.Write(record); err != nil {
		return "", err
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}
	return strings.TrimRight(buffer.String(), "\r\n"), nil
}

// Do not use this.
func FromCSV(k protoreflect.Kind, s string) (any, bool, error) {
	fn, ok := converters[k]
	if !ok {
		panic(fmt.Sprintf("unsupported type: %v", k))
	}

	reader := gocsv.NewReader(strings.NewReader(s))
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, false, err
	}
	if len(records) == 0 {
		// 型キャストでのpanicを防止するため、0要素でも型を付与
		return makeSlice(k, 0, 0), false, nil
	} else if len(records) > 1 {
		return nil, false, fmt.Errorf("1 or more csv records are not permitted: %v", len(records))
	}
	out := reflect.ValueOf(makeSlice(k, 0, len(records[0])))
	for _, s := range records[0] {
		v, err := fn(s)
		if err != nil {
			return nil, false, err
		}
		out = reflect.Append(out, reflect.ValueOf(v))
	}
	return out.Interface(), out.Len() > 0, nil
}

func makeSlice(k protoreflect.Kind, length, capacity int) any {
	switch k {
	case protoreflect.BoolKind:
		return make([]bool, length, capacity)
	case protoreflect.Int32Kind:
		return make([]int32, length, capacity)
	case protoreflect.Sint32Kind:
		return make([]int32, length, capacity)
	case protoreflect.Uint32Kind:
		return make([]uint32, length, capacity)
	case protoreflect.Int64Kind:
		return make([]int64, length, capacity)
	case protoreflect.Sint64Kind:
		return make([]int64, length, capacity)
	case protoreflect.Uint64Kind:
		return make([]uint64, length, capacity)
	case protoreflect.Sfixed32Kind:
		return make([]int32, length, capacity)
	case protoreflect.Fixed32Kind:
		return make([]int32, length, capacity)
	case protoreflect.FloatKind:
		return make([]float32, length, capacity)
	case protoreflect.Sfixed64Kind:
		return make([]int64, length, capacity)
	case protoreflect.Fixed64Kind:
		return make([]int64, length, capacity)
	case protoreflect.DoubleKind:
		return make([]float64, length, capacity)
	case protoreflect.StringKind:
		return make([]string, length, capacity)
	default:
		panic(fmt.Sprintf("unsupported type: %v", k))
	}
}

func convertToTime(s string) (time.Time, bool) {
	var out time.Time
	for _, l := range []string{
		"20060102",
		"2006/1/2",
		"2006-1-2",
	} {
		if t, err := time.ParseInLocation(l, s, tz); err != nil {
			continue
		} else {
			out = t
			break
		}
	}
	if out.IsZero() {
		return out, false
	}
	return out, true
}

var (
	converters = map[protoreflect.Kind]func(string) (any, error){
		protoreflect.BoolKind: func(s string) (any, error) {
			switch s {
			case "0", "":
				return false, nil
			case "1":
				return true, nil
			default:
				return false, &csv.Error{
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000017,
				}
			}
		},
		protoreflect.Int32Kind: func(s string) (any, error) {
			if s == "" {
				return int32(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return int32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return int32(v), nil
		},
		protoreflect.Sint32Kind: func(s string) (any, error) {
			if s == "" {
				return int32(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return int32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return int32(v), nil
		},
		protoreflect.Uint32Kind: func(s string) (any, error) {
			if s == "" {
				return uint32(0), nil
			}
			v, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				return uint32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return uint32(v), err
		},
		protoreflect.Int64Kind: func(s string) (any, error) {
			if s == "" {
				return int64(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return int64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return v, nil
		},
		protoreflect.Sint64Kind: func(s string) (any, error) {
			if s == "" {
				return int64(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return int64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return v, nil
		},
		protoreflect.Uint64Kind: func(s string) (any, error) {
			if s == "" {
				return uint64(0), nil
			}
			v, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return uint64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return uint64(v), err
		},
		protoreflect.Sfixed32Kind: func(s string) (any, error) {
			if s == "" {
				return int32(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return int32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return int32(v), err
		},
		protoreflect.Fixed32Kind: func(s string) (any, error) {
			if s == "" {
				return uint32(0), nil
			}
			v, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				return uint32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return uint32(v), err
		},
		protoreflect.FloatKind: func(s string) (any, error) {
			if s == "" {
				return float32(0), nil
			}
			v, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return float32(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return float32(v), err
		},
		protoreflect.Sfixed64Kind: func(s string) (any, error) {
			if s == "" {
				return int64(0), nil
			}
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return int64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return v, nil
		},
		protoreflect.Fixed64Kind: func(s string) (any, error) {
			if s == "" {
				return uint64(0), nil
			}
			v, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return uint64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return uint64(v), nil
		},
		protoreflect.DoubleKind: func(s string) (any, error) {
			if s == "" {
				return float64(0), nil
			}
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return float64(0), &csv.Error{
					Causes:      []error{err},
					ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000004,
				}
			}
			return v, nil
		},
		protoreflect.StringKind: func(s string) (any, error) {
			return s, nil
		},
	}
)
