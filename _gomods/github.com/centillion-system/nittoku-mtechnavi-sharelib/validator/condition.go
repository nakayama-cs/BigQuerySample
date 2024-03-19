package validator

import (
	"context"
)

var (
	// フィールド値が空か判定
	EmptyField = func(name string) Condition {
		return func(ctx context.Context, m *MessageWrapper) bool {
			fd := m.FieldDesc(name)
			fv := m.ProtoReflect().Get(fd)
			return !notEmpty(fd, fv)
		}
	}
	// フィールド値が空でないか判定
	NotEmptyField = func(name string) Condition {
		return func(ctx context.Context, m *MessageWrapper) bool {
			fd := m.FieldDesc(name)
			fv := m.ProtoReflect().Get(fd)
			return notEmpty(fd, fv)
		}
	}
)
