package ioengine

import (
	"bufio"
	"context"
	gocsv "encoding/csv"
	"errors"
	"fmt"
	"io"
	"mtechnavi/sharelib/constants"
	"mtechnavi/sharelib/dataproxy/codec"
	"mtechnavi/sharelib/encoding/csv"
	sharelibpb "mtechnavi/sharelib/protobuf"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	"mtechnavi/sharelib/text/encoding"
	"mtechnavi/sharelib/validator"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/go-multierror"
	enc "golang.org/x/text/encoding"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrIoEngineFileError   = errors.New("IoEngine:File Error")
	ErrIoEngineColumnError = errors.New("IoEngine:Column Error")
)

// インポート時のErrorCSV出力
func ExportErrorCSVToSignedURL(
	ctx context.Context,
	format *sharelibpb.FileFormat,
	readableSignedURL string,
	writableSignedURL string,
	importResults []*ImportResult,
	errorHeaderName string,
	warningHeaderName string,
) error {

	enc := toEncoding(format.EncodingName)

	writeBody, pw := io.Pipe()
	w := encoding.NewWriter(pw, enc)
	writer := bufio.NewWriter(w)

	splitChar := ","
	switch format.CommaName {
	case "COMMA":
		splitChar = ","
	case "TAB":
		splitChar = "\t"
	}

	readReq, err := http.NewRequestWithContext(ctx, http.MethodGet, readableSignedURL, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to NewRequestWithContext(ReadReq): %w", err)
	}

	readRes, err := http.DefaultClient.Do(readReq)
	if err != nil {
		return fmt.Errorf("failed to httpCli.Do: %w", err)
	}
	defer readRes.Body.Close()

	if readRes.StatusCode >= 400 {
		return errors.New(`httpRes.StatusCode >= 400`)
	}
	r := encoding.NewReader(readRes.Body, enc)
	scanner := bufio.NewScanner(r)

	errmap := map[int]*ImportResult{}
	for _, importResult := range importResults {
		es := importResult.AllErrors()
		var csvLine int
		// 行数を特定
		for _, e := range es {
			switch c := e.(type) {
			case *csv.Error:
				csvLine = c.Line
			}
		}
		errmap[csvLine] = importResult
	}

	go func() {
		count := 1
		if format.Header {
			line := ""
			if scanner.Scan() {
				line = scanner.Text()
				writeString := fmt.Sprintf("%s%s\"%s\"%s\"%s\"\n", line, splitChar, errorHeaderName, splitChar, warningHeaderName)
				if _, err := writer.WriteString(writeString); err != nil {
					pw.CloseWithError(err)
				}
				count++
			} else {
				if scanner.Err() != nil {
					pw.CloseWithError(err)
				}
			}
		}
		for {
			line := ""
			if scanner.Scan() {
				line = scanner.Text()

				var errorContent, warningContent string
				if result, ok := errmap[count]; ok {
					errorContent = strings.Join(result.ErrorContents("ja", format), ",")
					warningContent = strings.Join(result.WarningContents("ja", format), ",")
				}

				writeString := fmt.Sprintf("%s%s\"%s\"%s\"%s\"\n", line, splitChar, errorContent, splitChar, warningContent)
				if _, err := writer.WriteString(writeString); err != nil {
					pw.CloseWithError(err)
				}
				count++
			} else {
				if scanner.Err() != nil {
					pw.CloseWithError(err)
				}
				break
			}
		}
		writer.Flush()
		if err := pw.Close(); err != nil {
			pw.CloseWithError(err)
		}
	}()

	writeReq, err := http.NewRequestWithContext(ctx, http.MethodPut, writableSignedURL, writeBody)
	if err != nil {
		return fmt.Errorf("failed to NewRequestWithContext(WriteReq): %w", err)
	}
	writeReq.Header.Set("Content-Type", "text/csv")

	writeRes, err := http.DefaultClient.Do(writeReq)
	if err != nil {
		return fmt.Errorf("failed to http.DefaultClient.Do(WriteReq): %w", err)
	}
	defer writeRes.Body.Close()

	// レスポンスボディをcloseするために一度Readして読み切る
	if _, err := io.ReadAll(writeRes.Body); err != nil {
		return fmt.Errorf("failed to io.ReadAll: %w", err)
	}
	if writeRes.StatusCode >= 400 {
		return errors.New(`httpRes.StatusCode >= 400`)
	}

	return nil
}

func ExportCSVToSignedURL(
	ctx context.Context,
	format *sharelibpb.FileFormat,
	signedURL string,
	records []proto.Message,
	transformer func(proto.Message) csv.Marshaler,
) error {
	reqBody, pw := io.Pipe()

	enc := toEncoding(format.EncodingName)
	w := encoding.NewWriter(pw, enc)
	writer := csv.NewWriter(w)
	writer.Comma = toComma(format.CommaName)
	writer.ForceQuote = !format.QuoteMinimally
	writer.UseCRLF = true

	go func() {
		if format.Header {
			var l []string
			for _, hc := range format.GetHeaderColumns() {
				l = append(l, hc.DisplayName)
			}
			if err := writer.Write(csv.StringMarshaler(l)); err != nil {
				pw.CloseWithError(err)
			}
		}
		for _, record := range records {
			v := transformer(record)
			if err := writer.Write(v); err != nil {
				pw.CloseWithError(err)
			}
		}
		if err := pw.Close(); err != nil {
			pw.CloseWithError(err)
		}
	}()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, signedURL, reqBody)
	if err != nil {
		return fmt.Errorf("failed to NewRequestWithContext: %w", err)
	}
	httpReq.Header.Set("Content-Type", "text/csv")

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to http.DefaultClient.Do: %w", err)
	}
	defer httpRes.Body.Close()

	// レスポンスボディをcloseするために一度Readして読み切る
	if _, err := io.ReadAll(httpRes.Body); err != nil {
		return fmt.Errorf("failed to io.ReadAll: %w", err)
	}
	if httpRes.StatusCode >= 400 {
		return errors.New(`httpRes.StatusCode >= 400`)
	}

	return nil
}

type ImportResult struct {
	proto.Message
	Err error
}

type ImportResultError interface {
	GetLine() int
	GetColumn() int
	GetErrorContent(lang string) string
	GetErrorLevel() sharelibpb.ErrorLevel
}

func (r *ImportResult) ErrorContents(lang string, format *sharelibpb.FileFormat) []string {
	var msgs []string
	for _, e := range r.AllErrors() {
		v, ok := e.(ImportResultError)
		if !ok {
			continue
		}
		content := v.GetErrorContent(lang)
		columnName := format.HeaderColumns[v.GetColumn()-1].DisplayName
		if content != "" && v.GetErrorLevel() == sharelibpb.ErrorLevel_ERROR {
			msgs = append(msgs, fmt.Sprintf("[%s]%s", columnName, content))
		}
	}
	return msgs
}

func (r *ImportResult) WarningContents(lang string, format *sharelibpb.FileFormat) []string {
	var msgs []string
	for _, e := range r.AllErrors() {
		v, ok := e.(ImportResultError)
		if !ok {
			continue
		}
		content := v.GetErrorContent(lang)
		columnName := format.HeaderColumns[v.GetColumn()-1].DisplayName
		if content != "" && v.GetErrorLevel() == sharelibpb.ErrorLevel_WARNING {
			msgs = append(msgs, fmt.Sprintf("[%s]%s", columnName, content))
		}
	}
	return msgs
}

func (r *ImportResult) AllErrors() []error {
	var errs []error
	if r.Err == nil {
		return errs
	}
	switch e := r.Err.(type) {
	case *multierror.Error:
		errs = append(errs, e.WrappedErrors()...)
	default:
		errs = append(errs, e)
	}
	return errs
}

func ImportCSVFromSignedURL(
	ctx context.Context,
	format *sharelibpb.FileFormat,
	signedURL string,
	formatDefaults interface{},
	baseRecords interface{},
	transformer func(csv.Unmarshaler, proto.Message) (proto.Message, error),
	opts ...ImportOption,
) ([]*ImportResult, error) {
	elTyp, err := getElenemtTypeOfSlice(formatDefaults)
	if err != nil {
		return nil, fmt.Errorf("failed to getElementTypeOfSlice(formatDefaults): %w", err)
	}
	_, ok := reflect.New(elTyp).Elem().Interface().(csv.Unmarshaler)
	if !ok {
		return nil, fmt.Errorf("type mismatch: element type=%T must be csv.Unmarshaler", reflect.New(elTyp).Elem().Interface())
	}
	relTyp, err := getElenemtTypeOfSlice(baseRecords)
	if err != nil {
		return nil, fmt.Errorf("failed to getElementTypeOfSlice(baseRecords): %w", err)
	}
	rec, ok := reflect.New(relTyp).Elem().Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("type mismatch: element type=%T must be proto.Message", reflect.New(relTyp).Elem().Interface())
	}
	typeName := string(rec.ProtoReflect().Descriptor().FullName())

	params, err := GetImportParamsFromOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to GetImportParamsFromOptions: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, signedURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to NewRequestWithContext: %w", err)
	}
	httpCli := &http.Client{}
	httpRes, err := httpCli.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to httpCli.Do: %w", err)
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode >= 400 {
		return nil, errors.New(`httpRes.StatusCode >= 400`)
	}
	enc := toEncoding(format.EncodingName)
	r := encoding.NewReader(httpRes.Body, enc)
	reader := csv.NewReader(r)
	reader.Comma = toComma(format.CommaName)
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	var out []*ImportResult
	// header
	if format.Header {
		if err := reader.Read(&csv.HeaderUnmarshaler{}); isCSVFileError(err) {
			// 元の処理に合わせてCSVのエラーに関してのみ捕まえる
			return nil, ErrIoEngineFileError
		}
	}
	for {
		val := reflect.New(elTyp.Elem()).Interface().(csv.Unmarshaler)
		// read
		var result ImportResult
		if err := reader.Read(val); errors.Is(err, io.EOF) {
			break
		} else if isCSVFileError(err) {
			return nil, ErrIoEngineFileError
		} else if isCSVColumnError(err) {
			return nil, ErrIoEngineColumnError
		} else if err != nil {
			result.Err = multierror.Append(result.Err, err)
		}
		// validate
		if params.RuleSet != nil {
			if err := validate(ctx, params.RuleSet, val); err != nil {
				for _, e := range toCsvErrors(err) {
					if e.Column > 0 {
						e.Line, _ = reader.FieldPos(e.Column - 1)
					}
					result.Err = multierror.Append(result.Err, e)
				}
			}
		}
		// transform
		base := reflect.New(relTyp.Elem()).Interface().(proto.Message)
		if v, ok := val.(proto.Message); ok {
			m := messageWrapper{v}
			pkName := codec.GetPrimaryKeyName(typeName)
			if v := m.getBaseRecord(pkName, baseRecords); v != nil {
				base = v
			}
			if m, err := transformer(val, base); err != nil {
				return nil, fmt.Errorf("failed to transformer: %w", err)
			} else {
				result.Message = m
			}
		}
		out = append(out, &result)
	}
	return out, nil
}

func validate(ctx context.Context, rs *validator.RuleSet, val csv.Unmarshaler) error {
	v, ok := val.(proto.Message)
	if !ok {
		return nil
	}
	return rs.Eval(ctx, v)
}

type messageWrapper struct {
	proto.Message
}

func (m *messageWrapper) getBaseRecord(pkName string, baseRecords interface{}) proto.Message {
	if !m.NotEmptyPrimaryKey(pkName) {
		return nil
	}
	pkVal := m.Get(pkName).String()
	if v := getRecordByPrimaryKey(pkName, pkVal, baseRecords); v != nil {
		return v
	}
	return nil
}

func (m *messageWrapper) Get(name string) protoreflect.Value {
	fd := m.fieldDesc(name)
	if fd == nil {
		return protoreflect.Value{}
	}
	return m.ProtoReflect().Get(fd)
}

func (m *messageWrapper) fieldDesc(name string) protoreflect.FieldDescriptor {
	pr := m.ProtoReflect()
	fd := pr.Descriptor().Fields().ByTextName(name)
	return fd
}

func (m *messageWrapper) NotEmptyPrimaryKey(pkName string) bool {
	fd := m.fieldDesc(pkName)
	if fd == nil {
		return false
	}
	fv := m.ProtoReflect().Get(fd)
	if !fv.IsValid() {
		return false
	}
	rv := reflect.ValueOf(fv.Interface())
	return rv.IsValid() && !rv.IsZero()
}

func getRecordByPrimaryKey(pkName, pkVal string, records interface{}) proto.Message {
	rv := reflect.ValueOf(records)
	for i := 0; i < rv.Len(); i++ {
		rec := rv.Index(i)
		var pk string
		if v, ok := rec.Interface().(proto.Message); ok {
			m := messageWrapper{v}
			pk = m.Get(pkName).String()
		}
		if pk == pkVal {
			return rec.Interface().(proto.Message)
		}
	}
	return nil
}

func toComma(commaName string) rune {
	switch commaName {
	case "TAB":
		return '\t'
	case "COMMA":
		return ','
	default:
		return ','
	}
}

func toEncoding(encodingName string) enc.Encoding {
	switch encodingName {
	case "SHIFT_JIS":
		return encoding.ShiftJIS
	case "UTF_8":
		return encoding.UTF8
	default:
		return encoding.ShiftJIS
	}
}

func toCsvErrors(err error) []*csv.Error {
	var out []*csv.Error
	switch e := err.(type) {
	case *validator.Error:
		for _, v := range e.Results {
			csvErr := &csv.Error{
				ErrorLevel:  v.ErrorLevel,
				MessageName: v.MessageName,
				MessageArgs: v.MessageArgs,
			}
			if opt, ok := GetFieldExtension(v.FieldDesc()); ok {
				csvErr.Column = int(opt.Column)
			}
			out = append(out, csvErr)
		}
	}
	return out
}

func GetFieldExtension(fd protoreflect.FieldDescriptor) (*mtnpb.CsvColumnOptions, bool) {
	if fd == nil {
		return nil, false
	}
	var out *mtnpb.CsvColumnOptions
	proto.RangeExtensions(fd.Options(), func(xt protoreflect.ExtensionType, ev interface{}) bool {
		v, ok := ev.(*mtnpb.CsvColumnOptions)
		if !ok {
			return true
		}
		out = v
		return false
	})
	return out, out != nil
}

func getElenemtTypeOfSlice(v interface{}) (reflect.Type, error) {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("v is must be a pointer to a slice: %T", v)
	}
	elTyp := rv.Type().Elem()
	return elTyp, nil
}

func isCSVFileError(e error) bool {
	switch {
	case errors.Is(e, gocsv.ErrTrailingComma):
		fallthrough
	case errors.Is(e, gocsv.ErrBareQuote):
		fallthrough
	case errors.Is(e, gocsv.ErrQuote):
		fallthrough
	case errors.Is(e, gocsv.ErrFieldCount):
		return true
	default:
		return false
	}
}

func isCSVColumnError(err error) bool {
	var e *csv.Error
	if errors.As(err, &e) {
		if e.MessageName == constants.MessageName_E0000022 {
			return true
		}
	}
	return false
}
