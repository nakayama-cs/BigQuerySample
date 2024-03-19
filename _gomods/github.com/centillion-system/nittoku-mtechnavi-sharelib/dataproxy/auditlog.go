package dataproxy

import (
	"context"
)

func (c *Client) ListAuditLogs(ctx context.Context, options ...Option) error {
	// TODO 実装
	panic("unimplemented yet")
	// c.c.ListAuditLogs(ctx, &dataproxypb.ListAuditLogsRequest{
	// 	TenantId:         "",
	// 	UserId:           "",
	// 	AuditCategory:    "",
	// 	AuditSubCategory: "",
	// 	Table:            "",
	// 	RecordId:         "",
	// 	From:             0,
	// 	To:               0,
	// 	Limit:            0,
	// 	PageToken:        "",
	// })
}
