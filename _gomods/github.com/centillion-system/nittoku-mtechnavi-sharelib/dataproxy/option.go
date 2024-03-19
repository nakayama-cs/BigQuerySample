package dataproxy

import (
	"fmt"
	dataproxypb "mtechnavi/dataproxy/protobuf"
)

type writeOptions struct {
}

type getOptions struct {
}

type listOptions struct {
	TenantId string

	PageToken string

	RecordIds []string

	Limit int

	IndexFilter *indexFilter

	Filter string
}

type indexFilter struct {
	Name string

	// One of:
	// - "==", "!="
	// - ">", ">=", "<=", "<"
	// - "in", "not-in"
	Op string

	Value any
}

type shareOptions struct {
	Mode dataproxypb.ShareToMode
}

type Option func(any)

func WithTenantId(v string) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.TenantId = v
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithPageToken(v string) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.PageToken = v
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithRecordIds(v []string) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.RecordIds = v
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithLimit(v int) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.Limit = v
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithIndexFilter(name, op string, value any) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.IndexFilter = &indexFilter{
				Name:  name,
				Op:    op,
				Value: value,
			}
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithFilter(v string) Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *listOptions:
			o.Filter = v
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}

func WithShareToModeIncremental() Option {
	return func(opts any) {
		switch o := opts.(type) {
		case *shareOptions:
			o.Mode = dataproxypb.ShareToMode_Incremental
		default:
			panic(fmt.Sprintf("unsupported with %T", opts))
		}
	}
}
