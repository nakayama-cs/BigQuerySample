package validator

import (
	"context"
	"fmt"
	"mtechnavi/sharelib/auth"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/protobuf"
	sharelibpb "mtechnavi/sharelib/protobuf"
	"reflect"
	"strings"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Error struct {
	Results []*RuleResult
}

func (e *Error) Error() string {
	var msgs []string
	for _, v := range e.Results {
		msg := fmt.Sprintf("RuleResult Field:%s,ErrorLevel:%s,MessageName:%s,MessageArgs:%v", v.Field, v.ErrorLevel, v.MessageName, v.MessageArgs)
		msgs = append(msgs, msg)
	}

	return strings.Join(msgs, ",")
}

func (e *Error) GRPCStatus() *status.Status {
	return status.New(codes.Unimplemented, "")
}

type RuleSet struct {
	// 評価済みのメッセージ一覧
	messages []*MessageWrapper

	// バリデートルール一覧
	rules []Rule
}

func NewRuleSet() *RuleSet {
	return &RuleSet{}
}

func (rs *RuleSet) Add(rules ...Rule) {
	rs.rules = append(rs.rules, rules...)
}

func (rs *RuleSet) Eval(ctx context.Context, m proto.Message) error {
	wrapper := &MessageWrapper{
		Message: m,
	}
	var l []*RuleResult
	for _, rule := range rs.rules {
		if rs := rule.Eval(ctx, wrapper); rs != nil {
			for _, r := range rs {
				r.message = wrapper
				l = append(l, r)
			}
		}
	}
	rs.messages = append(rs.messages, wrapper)
	if len(l) == 0 {
		return nil
	}
	return &Error{
		Results: l,
	}
}

type Rule interface {
	Eval(context.Context, *MessageWrapper) []*RuleResult
}

// Ruleインターフェースを実装したルール関数
type RuleFunc func(context.Context, *MessageWrapper) []*RuleResult

func (fn RuleFunc) Eval(ctx context.Context, m *MessageWrapper) []*RuleResult {
	return fn(ctx, m)
}

type RuleResult struct {
	message *MessageWrapper

	// バリデート対象となるフィールド名を指定する
	//
	// フィールド名を指定しない場合は、
	// レコード単体の不正フィールドエラーではなく、システム的なエラーと判定する
	Field string

	// エラー種別
	ErrorLevel pb.ErrorLevel

	// システムで定義されたメッセージコード
	MessageName constants.MessageName

	// メッセージパラメータ
	MessageArgs []any
}

type MessageWrapper struct {
	proto.Message
}

func (r *RuleResult) FieldDesc() protoreflect.FieldDescriptor {
	if r.Field == "" {
		return nil
	}
	return r.message.FieldDesc(r.Field)
}

func (m *MessageWrapper) Get(field string) protoreflect.Value {
	fd := m.FieldDesc(field)
	return m.ProtoReflect().Get(fd)
}

func (m *MessageWrapper) FieldDesc(field string) protoreflect.FieldDescriptor {
	pr := m.ProtoReflect()
	fd := pr.Descriptor().Fields().ByTextName(field)
	if fd == nil {
		panic(fmt.Sprintf("unexpected field:%s", field))
	}
	return fd
}

func NotEmptyf(field string, level pb.ErrorLevel, messageName constants.MessageName, args ...any) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		fd := m.FieldDesc(field)
		fv := m.ProtoReflect().Get(fd)
		if !notEmpty(fd, fv) {
			return []*RuleResult{
				{
					Field:       field,
					ErrorLevel:  level,
					MessageName: messageName,
					MessageArgs: args,
				},
			}
		}
		return nil
	})
}

func GreaterThanOrEqual(field string, min int64) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		fd := m.FieldDesc(field)
		fv := m.ProtoReflect().Get(fd)
		if fd.Cardinality() == protoreflect.Repeated {
			if fd.IsList() {
				for i := 0; i < fv.List().Len(); i++ {
					if r := isGte(field, fd, fv, min); r != nil {
						return r
					}
				}
			} else {
				panic(fmt.Sprintf("unsupported field: type=%T", fv.Interface()))
			}
		}
		return isGte(field, fd, fv, min)
	})
}

func LessThanOrEqual(field string, max int64) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		fd := m.FieldDesc(field)
		fv := m.ProtoReflect().Get(fd)
		if fd.Cardinality() == protoreflect.Repeated {
			if fd.IsList() {
				for i := 0; i < fv.List().Len(); i++ {
					if r := isLte(field, fd, fv, max); r != nil {
						return r
					}
				}
			} else {
				panic(fmt.Sprintf("unsupported field: type=%T", fv.Interface()))
			}
		}
		return isLte(field, fd, fv, max)
	})
}

func ContainsAnyf(field string, l []string, format string, args ...any) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		v := m.Get(field).String()
		for i := range l {
			if l[i] == v {
				return []*RuleResult{
					{
						Field:       field,
						MessageArgs: args,
					},
				}
			}
		}
		return nil
	})
}

func (rs *RuleSet) Unique(field string, level pb.ErrorLevel, messageName constants.MessageName) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		v := m.Get(field).String()
		for _, rm := range rs.messages {
			mv := rm.Get(field).String()
			if mv == "" {
				continue
			}
			if mv == v {
				return []*RuleResult{
					{
						Field:       field,
						ErrorLevel:  level,
						MessageName: messageName,
					},
				}
			}
		}
		return nil
	})
}

type Condition func(ctx context.Context, m *MessageWrapper) bool

func (rs *RuleSet) EvalIf(expr Condition, rules ...Rule) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		if !expr(ctx, m) {
			return nil
		}
		var l []*RuleResult
		for _, r := range rules {
			if res := r.Eval(ctx, m); res != nil {
				l = append(l, res...)
			}
		}
		return l
	})
}

func isGte(field string, fd protoreflect.FieldDescriptor, fv protoreflect.Value, min int64) []*RuleResult {
	switch fd.Kind() {
	case protoreflect.StringKind:
		if int64(utf8.RuneCountInString(fv.String())) < min {
			return []*RuleResult{
				{
					Field:       field,
					MessageName: constants.MessageName_E0000011,
					MessageArgs: []any{min},
				},
			}
		}
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		if fv.Int() < min {
			return []*RuleResult{
				{
					Field:       field,
					MessageName: constants.MessageName_E0000013,
					MessageArgs: []any{min},
				},
			}
		}
	default:
		panic(fmt.Sprintf("unsupported field: type=%T", fv.Interface()))
	}
	return nil
}

func isLte(field string, fd protoreflect.FieldDescriptor, fv protoreflect.Value, max int64) []*RuleResult {
	switch fd.Kind() {
	case protoreflect.StringKind:
		if int64(utf8.RuneCountInString(fv.String())) > max {
			return []*RuleResult{
				{
					Field:       field,
					ErrorLevel:  pb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000012,
					MessageArgs: []any{max},
				},
			}
		}
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		if fv.Int() > max {
			return []*RuleResult{
				{
					Field:       field,
					ErrorLevel:  pb.ErrorLevel_ERROR,
					MessageName: constants.MessageName_E0000014,
					MessageArgs: []any{max},
				},
			}
		}
	default:
		panic(fmt.Sprintf("unsupported field: type=%T", fv.Interface()))
	}
	return nil
}

func notEmpty(fd protoreflect.FieldDescriptor, fv protoreflect.Value) bool {
	if fd.Cardinality() == protoreflect.Repeated {
		if fd.IsList() {
			return fv.List().Len() > 0
		} else if fd.IsMap() {
			return fv.Map().Len() > 0
		} else {
			panic(fmt.Sprintf("unsupported field: type=%T", fv.Interface()))
		}
	} else {
		rv := reflect.ValueOf(fv.Interface())
		if rv.Kind() == reflect.Pointer {
			rv = rv.Elem()
		}
		return rv.IsValid() && !rv.IsZero()
	}
}

type ValidationFunc func(context.Context, proto.Message, []*proto.Message, ...string) (bool, string, []any, error)

func GetCSVFieldName(baseName string) string {
	// xxxx.{{FieldName}}のため、「.」でsplitし、最後の文字列を取得
	s := strings.Split(baseName, ".")
	f := s[len(s)-1]

	return f
}

// 取り込みバリデーション
func (rs *RuleSet) ImportValidate(fn ValidationFunc, level pb.ErrorLevel, message constants.MessageName, field string, max int64, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, level, message, fn, field, fields...)
	})
}

// 必須入力(E0000003)
func (rs *RuleSet) E0000003(fn ValidationFunc, field string, fields ...string) Rule {
	if fn != nil {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000003, fn, field, fields...)
		})
	} else {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			splits := strings.Split(field, ".")
			trimField := splits[len(splits)-1]
			fd := m.FieldDesc(trimField)
			fv := m.ProtoReflect().Get(fd)

			if fv.String() == "" {
				return []*RuleResult{
					{
						Field:       field,
						ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
						MessageName: constants.MessageName_E0000003,
						MessageArgs: nil,
					},
				}
			}
			return nil
		})
	}
}

// X文字以下(E0000012)
func (rs *RuleSet) E0000012(fn ValidationFunc, field string, max int64, fields ...string) Rule {
	if fn != nil {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000012, fn, field, fields...)
		})
	} else {
		return LessThanOrEqual(field, max)
	}
}

// 取引先企業ステータス=取引ありのみ許可(E0000033)
func (rs *RuleSet) E0000033(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000033, fn, field, fields...)
	})
}

// 取引先管理場所IDで取得したレコードの取引窓口＝ONのみ許可(E0000034)
func (rs *RuleSet) E0000034(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000034, fn, field, fields...)
	})
}

// 図番と枝番で図面マスタ存在のみ許可(E0000036)
// 更新番号が1以上時は図番と枝番と更新番号で図面マスタ存在のみ許可(E0000036)
func (rs *RuleSet) E0000036(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000036, fn, field, fields...)
	})
}

// 特権管理グループに対するユーザー追加/除外は権限がある場合のみ許可(E0000057)
func (rs *RuleSet) E0000057(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000057, fn, field, fields...)
	})
}

// 自身の操作で自ユーザーを特権管理グループから除外することは不可(E0000058)
func (rs *RuleSet) E0000058(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E0000058, fn, field, fields...)
	})
}

// 未入力 (必須入力) (E1000001)
func (rs *RuleSet) E1000001(fn ValidationFunc, field string, fields ...string) Rule {
	if fn != nil {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000001, fn, field, fields...)
		})
	} else {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			splits := strings.Split(field, ".")
			trimField := splits[len(splits)-1]
			fd := m.FieldDesc(trimField)
			fv := m.ProtoReflect().Get(fd)

			if fv.String() == "" {
				return []*RuleResult{
					{
						Field:       field,
						ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
						MessageName: constants.MessageName_E1000001,
						MessageArgs: nil,
					},
				}
			}
			return nil
		})
	}
}

// 新規モード時は特定値のみ許可(E1000002)
func (rs *RuleSet) E1000002(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000002, fn, field, fields...)
	})
}

// 名称マスタに存在のみ許可(E1000003)
func (rs *RuleSet) E1000003(
	lookupByCode func(context.Context, string, string, string) (*sharelibpb.NameOption, error),
	field string,
	categoryName string,
	itemName string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		invoker := auth.InvokerFromContext(ctx)

		splits := strings.Split(field, ".")
		trimField := splits[len(splits)-1]
		fd := m.FieldDesc(trimField)
		fv := m.ProtoReflect().Get(fd)

		codes := []string{}
		if fd.IsList() {
			l := m.ProtoReflect().Get(fd).List()
			for i := 0; i < l.Len(); i++ {
				li := l.Get(i).String()
				codes = append(codes, li)
			}
		} else {
			codes = append(codes, fv.String())
		}

		for _, code := range codes {
			if code == "" {
				// コードが未指定時はvalidationの対象外
				continue
			}

			if _, err := lookupByCode(ctx, invoker.Tenant.TenantId, categoryName, code); err != nil {
				return []*RuleResult{
					{
						Field:       field,
						ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
						MessageName: constants.MessageName_E1000003,
						MessageArgs: []any{itemName},
					},
				}
			}
		}
		return nil
	})
}

// 取込ファイル内の重複コードは不可(E1000005)
func (rs *RuleSet) E1000005(fn ValidationFunc, field string, fields ...string) Rule {
	if fn != nil {
		return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
			return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000005, fn, field, fields...)
		})
	} else {
		return rs.Unique(field, pb.ErrorLevel_ERROR, constants.MessageName_E1000005)
	}
}

// 新規モード時はマスタ内に重複コードは不可(E1000006)
func (rs *RuleSet) E1000006(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000006, fn, field, fields...)
	})
}

// 更新モード時はIDのレコードが存在時のみ許可(E1000007)
func (rs *RuleSet) E1000007(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000007, fn, field, fields...)
	})
}

// 新規モード時はユーザ名称マスタ内に種別コード＋システムコードのレコードが存在する場合のみ許可(E1000009)
func (rs *RuleSet) E1000009(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_ERROR, constants.MessageName_E1000009, fn, field, fields...)
	})
}

// マスタをコードで取得し存在しない場合は警告(W0000003)
func (rs *RuleSet) W0000003(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_WARNING, constants.MessageName_W0000003, fn, field, fields...)
	})
}

// 管理場所マスタを管理場所コードで取得し複数存在する場合は警告(W0000004)
func (rs *RuleSet) W0000004(fn ValidationFunc, field string, fields ...string) Rule {
	return RuleFunc(func(ctx context.Context, m *MessageWrapper) []*RuleResult {
		return validateField(ctx, m, rs.messages, pb.ErrorLevel_WARNING, constants.MessageName_W0000004, fn, field, fields...)
	})
}

func validateField(ctx context.Context, m *MessageWrapper, ms []*MessageWrapper, level pb.ErrorLevel, messageName constants.MessageName, fn ValidationFunc, field string, fields ...string) []*RuleResult {

	doneMessages := []*proto.Message{}
	for _, v := range ms {
		doneMessages = append(doneMessages, &v.Message)
	}

	ok, errField, args, err := fn(ctx, m.Message, doneMessages, fields...)
	switch {
	case err != nil:
		return []*RuleResult{
			{
				Field:       errField,
				ErrorLevel:  sharelibpb.ErrorLevel_ERROR,
				MessageName: constants.MessageName_E0000001,
				MessageArgs: []any{err},
			},
		}
	case !ok:
		return []*RuleResult{
			{
				Field:       errField,
				ErrorLevel:  level,
				MessageName: messageName,
				MessageArgs: args,
			},
		}
	default:
		return nil
	}
}
