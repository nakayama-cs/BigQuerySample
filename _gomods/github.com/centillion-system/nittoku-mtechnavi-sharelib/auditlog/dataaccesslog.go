package auditlog

import (
	"context"
	pb "mtechnavi/sharelib/auditlog/protobuf"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type LogParameters struct {
	UserId   string
	TenantId string
}

func UnaryServerDataAccessLogger(fn func(context.Context) LogParameters) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		params := fn(ctx)
		res, err := handler(ctx, req)
		if err != nil {
			return res, err
		}
		PrintDataAccessLog(ctx, params.TenantId, params.UserId, res)
		return res, err
	}
}

type record interface {
	GetRecordId() string
	GetTypeName() string
}

// データ(閲覧)アクセスログ
func PrintDataAccessLog(ctx context.Context, tenantId, userId string, res interface{}) {
	v, ok := res.(proto.Message)
	if !ok {
		return
	}
	rv := v.ProtoReflect()
	md := rv.Descriptor()
	fds := md.Fields()

	var datas []proto.Message
	if _, ok := rv.Interface().(record); ok {
		datas = append(datas, rv.Interface())
	} else {
		if fd := fds.ByName("items"); fd != nil {
			if !fd.IsList() {
				return
			}
			items := rv.Get(fd).List()
			for i := 0; i < items.Len(); i++ {
				item := items.Get(i)
				datas = append(datas, item.Message().Interface())
			}
		}
	}
	for _, d := range datas {
		if r, ok := d.(record); ok {
			writeDataAccessLog(ctx, tenantId, userId, "ACCESS_LOG", pb.MessageCode_DATA_ACCESSED, r)
		}
	}
}

func writeDataAccessLog(ctx context.Context, tenantId, userId, auditSubcategory string, message pb.MessageCode, record record) {
	writeLog(ctx, logRequest{
		EventID:          uuid.New().String(),
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_DATA_ACCESS.String(),
		AuditSubcategory: auditSubcategory,
		RecordID:         record.GetRecordId(),
		Table:            record.GetTypeName(),
		Message:          message,
		Columns:          strings.Join(getColumnsArray(record), ","),
	})
}

type field interface {
	GetName() string
	GetValue() []byte
}

func getColumnsArray(record record) []string {
	var columns []string
	r, ok := record.(proto.Message)
	if ok {
		rv := r.ProtoReflect()
		md := rv.Descriptor()
		fds := md.Fields()
		// fields配下のカラム名を取得
		if fd := fds.ByName("fields"); fd != nil {
			if !fd.IsMap() {
				return columns
			}
			rv.Get(fd).Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				m := v.Message()
				if !m.IsValid() {
					return false
				}
				if f, ok := m.Interface().(field); ok {
					columns = append(columns, f.GetName())
				}
				return true
			})
		}
	}
	return columns
}
