package to

import (
	"fmt"
	"mtechnavi/sharelib/encoding/csv"
	"mtechnavi/sharelib/ioengine"
	pb "mtechnavi/sharelib/protobuf"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"

	"google.golang.org/protobuf/proto"
)

func MergeFileFormat(src interface {
	proto.Message
	GetCsvFileOptions() (*mtnpb.CsvFileOptions, bool)
	GetCsvHeaderColumns() []string
}, mergeFormat *pb.FileFormat) *pb.FileFormat {
	out := ToFileFormat(src)

	displayNameMap := map[string]string{}
	for _, v := range mergeFormat.HeaderColumns {
		displayNameMap[v.MessageName] = v.DisplayName
	}

	for _, o := range out.HeaderColumns {
		if v, ok := displayNameMap[o.MessageName]; ok {
			if v != "" {
				o.DisplayName = v
			}
		}
	}
	out.CommaName = mergeFormat.CommaName
	out.EncodingName = mergeFormat.EncodingName
	out.Header = mergeFormat.Header

	return out
}

func ToFileFormat(src interface {
	proto.Message
	GetCsvFileOptions() (*mtnpb.CsvFileOptions, bool)
	GetCsvHeaderColumns() []string
}) *pb.FileFormat {
	var format pb.FileFormat
	l := src.GetCsvHeaderColumns()
	for _, v := range l {
		format.HeaderColumns = append(format.HeaderColumns, &pb.HeaderColumn{
			DisplayName: v,
			MessageName: v,
		})
	}
	if opts, ok := src.GetCsvFileOptions(); ok {
		format.Header = opts.Header
		format.CommaName = opts.Comma.String()
		format.EncodingName = opts.Encoding.String()
		format.QuoteMinimally = opts.QuoteMinimally
		format.TypeName = string(proto.MessageName(src))
	} else {
		format.Header = opts.Header
		format.CommaName = opts.Comma.String()
		format.EncodingName = toEncodingName(opts.Encoding)
		format.QuoteMinimally = opts.QuoteMinimally
		format.TypeName = string(proto.MessageName(src))
	}
	return &format
}

func ToImportErrors(src interface {
	AllErrors() []error
}) []*pb.ImportError {
	var out []*pb.ImportError
	for _, e := range src.AllErrors() {
		var err = pb.ImportError{
			ErrorLevel: pb.ErrorLevel_ERROR,
		}
		switch e := e.(type) {
		case *csv.Error:
			err.RowNumber = int32(e.Line)
			err.ColumnNumber = int32(e.Column)
			err.ErrorLevel = e.ErrorLevel
			err.MessageName = e.MessageName.String()
			err.MessageArguments = toStrings(e.MessageArgs)
		default:
			err.MessageName = e.Error()
		}
		out = append(out, &err)
	}
	return out
}

func ToImportSummary(results []*ioengine.ImportResult) *pb.ImportSummary {
	var out pb.ImportSummary
	var success int64
	for _, v := range results {
		if v.Err == nil {
			success++
		} else {
			isError := false
			for _, e := range v.AllErrors() {
				v, ok := e.(ioengine.ImportResultError)
				if !ok {
					continue
				}
				if v.GetErrorLevel() == pb.ErrorLevel_WARNING {
					// WARNINGは取込としては成功扱い
					continue
				}
				isError = true
				break
			}
			if !isError {
				success++
			}
		}
	}
	total := int64(len(results))
	out.Total = total
	out.Success = success
	out.Error = total - success
	return &out
}

func toStrings(args []any) []string {
	var out []string
	for _, v := range args {
		out = append(out, fmt.Sprint(v))
	}
	return out
}

func toEncodingName(enc mtnpb.FileEncoding) string {
	switch enc {
	case mtnpb.FileEncoding_SHIFT_JIS:
		return "SHIFT_JIS"
	case mtnpb.FileEncoding_UTF_8:
		return "UTF_8"
	default:
		return "SHIFT_JIS"
	}
}
