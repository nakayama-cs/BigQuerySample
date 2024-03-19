package dataproxy

import (
	"context"
	"encoding/json"
	"fmt"
	dataproxypb "mtechnavi/dataproxy/protobuf"
	"mtechnavi/sharelib/dataproxy/transportcodec"
	recordpb "mtechnavi/sharelib/protobuf/record"
	"reflect"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/protobuf/proto"
)

func (c *Client) CreateMessage(ctx context.Context, l ...proto.Message) error {
	return c.CreateMessageWithOptions(ctx, l)
}

func (c *Client) CreateMessageWithOptions(ctx context.Context, l []proto.Message, options ...Option) error {
	var opts writeOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode(l)
	if err != nil {
		return err
	}
	res, err := c.c.CreateRecords(ctx, &dataproxypb.CreateRecordsRequest{
		Records: records,
	})
	if err != nil {
		return err
	}
	return Decode(l, res.Items)
}

func (c *Client) CreateOrUpdateMessage(ctx context.Context, l ...proto.Message) error {
	return c.CreateOrUpdateMessageWithOptions(ctx, l)
}

func (c *Client) CreateOrUpdateMessageWithOptions(ctx context.Context, l []proto.Message, options ...Option) error {
	var opts writeOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode(l)
	if err != nil {
		return err
	}
	res, err := c.c.CreateOrUpdateRecords(ctx, &dataproxypb.CreateOrUpdateRecordsRequest{
		Records: records,
	})
	if err != nil {
		return err
	}
	return Decode(l, res.Items)
}

func (c *Client) DeleteMessage(ctx context.Context, l ...proto.Message) error {
	return c.DeleteMessageWithOptions(ctx, l)
}

func (c *Client) DeleteMessageWithOptions(ctx context.Context, l []proto.Message, options ...Option) error {
	var opts writeOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode(l)
	if err != nil {
		return err
	}
	_, err = c.c.DeleteRecords(ctx, &dataproxypb.DeleteRecordsRequest{
		Records: records,
	})
	return err
}

func (c *Client) GetMessage(ctx context.Context, m proto.Message, options ...Option) error {
	var opts getOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode([]proto.Message{m})
	if err != nil {
		return err
	}
	res, err := c.c.GetRecord(ctx, &dataproxypb.GetRecordRequest{
		Record: records[0],
	})
	if err != nil {
		return err
	}
	return Decode([]proto.Message{m}, []*recordpb.Record{res})
}

var (
	indexOps = map[string]dataproxypb.Operator{
		"==":     dataproxypb.Operator_EQ,
		"!=":     dataproxypb.Operator_NEQ,
		">":      dataproxypb.Operator_GT,
		">=":     dataproxypb.Operator_GTE,
		"<":      dataproxypb.Operator_LT,
		"<=":     dataproxypb.Operator_LTE,
		"in":     dataproxypb.Operator_IN,
		"not-in": dataproxypb.Operator_NIN,
	}
)

type ListResult struct {
	NextPageToken string
}

func assumeElementMessage(l any) proto.Message {
	// *[]*pb.X or []*pb.X
	typ := reflect.TypeOf(l)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf("l must be a pointer to slice or slice: %T", l))
	}
	// *pb.X
	typ = typ.Elem()
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	m, ok := reflect.New(typ).Interface().(proto.Message)
	if !ok {
		panic(fmt.Sprintf("l must be a pointer to slice or slice of proto.Message: %T", l))
	}
	return m
}

func (c *Client) ListMessageAll(ctx context.Context, out any, capacityPerRead int, options ...Option) (*ListResult, error) {
	logger := ctxlogrus.Extract(ctx)

	// check type of out
	outType := reflect.TypeOf(out)
	if outType.Kind() != reflect.Pointer {
		logger.Panicf("out must be a pointer to slice: %T", out)
	}
	outElemType := outType.Elem()
	if outElemType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("out must be a pointer to slice: %T", out))
	}

	outValue := reflect.ValueOf(out)
	emptySlice := reflect.MakeSlice(outElemType, 0, outValue.Elem().Cap())

	listMessages := func(curCapacityPerRead int) error {
		if curCapacityPerRead <= 0 {
			return fmt.Errorf("capacityPerRead must be greater than 0: %d", curCapacityPerRead)
		}
		overrideOptions := append(options, WithLimit(curCapacityPerRead))

		// clear all items of out
		outValue.Elem().Set(emptySlice)

		for {
			res, err := c.ListMessages(ctx, out, overrideOptions...)
			if err != nil {
				return err
			}

			if res.NextPageToken == "" {
				break
			}
			overrideOptions = append(options, WithPageToken(res.NextPageToken), WithLimit(curCapacityPerRead))
		}

		return nil

	}

	needRetry := func(err error) bool {
		return err != nil && strings.Contains(err.Error(), "received message larger than max")
	}

	curCapacityPerRead := capacityPerRead
	const maxRetry = 2
	for retry := 0; retry < maxRetry; retry++ {
		err := listMessages(curCapacityPerRead)
		if err != nil {
			if !needRetry(err) {
				return nil, err
			}

			logger.WithError(err).Errorf("occured 'larger than max' error by %T: capacityPerRead=%d", out, curCapacityPerRead)

			if maxRetry == retry+1 {
				// give up when consumed retry counts.
				return nil, err
			}
			curCapacityPerRead /= 2
		} else {
			break
		}
	}

	return &ListResult{}, nil
}

func (c *Client) ListMessages(ctx context.Context, out any, options ...Option) (*ListResult, error) {
	logger := ctxlogrus.Extract(ctx)

	if reflect.TypeOf(out).Kind() != reflect.Pointer {
		logger.Panicf("out must be a pointer to slice: %T", out)
	}
	m := assumeElementMessage(out)

	var opts listOptions
	for _, opt := range options {
		opt(&opts)
	}

	req := &dataproxypb.ListRecordsRequest{
		TypeName:  string(proto.MessageName(m)),
		Limit:     int32(opts.Limit),
		PageToken: opts.PageToken,
		TenantId:  opts.TenantId,
		RecordIds: opts.RecordIds,
		Filter:    opts.Filter,
	}
	if opts.IndexFilter != nil {
		op, ok := indexOps[opts.IndexFilter.Op]
		if !ok {
			logger.Panicf("invalid index filter operator: %q", opts.IndexFilter.Op)
		}
		value, err := json.Marshal(opts.IndexFilter.Value)
		if err != nil {
			logger.WithError(err).Panicf("invalid index filter value")
		}
		req.UseIndex = &dataproxypb.IndexFilter{
			Name:     opts.IndexFilter.Name,
			Operator: op,
			Value:    string(value),
		}
	}
	res, err := c.listMessages(ctx, req)
	if err != nil {
		return nil, err
	}
	rv := reflect.Indirect(reflect.ValueOf(out))
	for _, v := range res.Items {
		val, err := transportcodec.Decode(v)
		if err != nil {
			return nil, err
		}
		rv.Set(reflect.Append(rv, reflect.ValueOf(val)))
	}
	return &ListResult{
		NextPageToken: res.NextPageToken,
	}, nil
}

func (c *Client) listMessages(ctx context.Context, req *dataproxypb.ListRecordsRequest) (*dataproxypb.ListRecordsResponse, error) {
	if req.Limit > 0 || req.PageToken != "" {
		return c.c.ListRecords(ctx, req)
	}

	var out dataproxypb.ListRecordsResponse
	for {
		req.Limit = 1000
		res, err := c.c.ListRecords(ctx, req)
		if err != nil {
			return nil, err
		}
		out.Items = append(out.Items, res.Items...)
		if res.NextPageToken == "" {
			return &out, nil
		}
		req.PageToken = res.NextPageToken
	}
}

func (c *Client) UpdateMessage(ctx context.Context, l ...proto.Message) error {
	return c.UpdateMessageWithOptions(ctx, l)
}

func (c *Client) UpdateMessageWithOptions(ctx context.Context, l []proto.Message, options ...Option) error {
	var opts writeOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode(l)
	if err != nil {
		return err
	}
	res, err := c.c.UpdateRecords(ctx, &dataproxypb.UpdateRecordsRequest{
		Records: records,
	})
	if err != nil {
		return err
	}
	return Decode(l, res.Items)
}

func (c *Client) ShareMessage(ctx context.Context, toTenantIds []string, l ...proto.Message) error {
	return c.ShareMessageWithOptions(ctx, toTenantIds, l)
}

func (c *Client) ShareMessageWithOptions(ctx context.Context, toTenantIds []string, l []proto.Message, options ...Option) error {
	var opts shareOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode(l)
	if err != nil {
		return err
	}
	_, err = c.c.ShareRecord(ctx, &dataproxypb.ShareRecordRequest{
		ShareTo: toTenantIds,
		Mode:    opts.Mode,
		Records: records,
	})
	return err
}

func (c *Client) GetShareMessage(ctx context.Context, m proto.Message, options ...Option) error {
	var opts getOptions
	for _, opt := range options {
		opt(&opts)
	}

	records, err := Encode([]proto.Message{m})
	if err != nil {
		return err
	}
	res, err := c.c.GetShareRecord(ctx, &dataproxypb.GetShareRecordRequest{
		Record: records[0],
	})
	if err != nil {
		return err
	}
	return Decode([]proto.Message{m}, []*recordpb.Record{res})
}

type ShareTenant struct {
	*dataproxypb.ShareTenant
}

func (v *ShareTenant) ShareStatus() dataproxypb.SHARE_STATUS {
	st := v.ShareTenant.ShareStatus
	val, ok := dataproxypb.SHARE_STATUS_value[st]
	if !ok {
		return dataproxypb.SHARE_STATUS_UNKNOWN
	}
	return dataproxypb.SHARE_STATUS(val)
}

func (c *Client) ListSharedTenants(ctx context.Context, m proto.Message, filterTenantIds []string, options ...Option) ([]*ShareTenant, error) {
	// TODO with options

	records, err := Encode([]proto.Message{m})
	if err != nil {
		return nil, err
	}
	res, err := c.c.ListSharedTenants(ctx, &dataproxypb.ListSharedTenantsRequest{
		Record:    records[0],
		TenantIds: filterTenantIds,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*ShareTenant, len(res.ShareTenants))
	for i := range res.ShareTenants {
		out[i] = &ShareTenant{
			ShareTenant: res.ShareTenants[i],
		}
	}
	return out, nil
}

func (c *Client) UnshareMessage(ctx context.Context, toTenantIds []string, l ...proto.Message) error {
	records, err := Encode(l)
	if err != nil {
		return err
	}
	_, err = c.c.UnshareRecord(ctx, &dataproxypb.UnshareRecordRequest{
		UnshareTo: toTenantIds,
		Records:   records,
	})
	return err
}
