package dataproxy

import (
	"context"
	dataproxypb "mtechnavi/dataproxy/protobuf"
)

func (c *Client) CreateAssetData(ctx context.Context, urlStr string, options ...Option) (*dataproxypb.AssetData, error) {
	// TODO with options

	return c.c.CreateAssetData(ctx, &dataproxypb.CreateAssetDataRequest{
		AssetData: &dataproxypb.AssetData{
			SourceUrl: urlStr,
		},
	})
}

func (c *Client) GetAssetData(ctx context.Context, assetDataId string, options ...Option) (*dataproxypb.AssetData, error) {
	// TODO with options

	return c.c.GetAssetData(ctx, &dataproxypb.GetAssetDataRequest{
		AssetDataId: assetDataId,
	})
}

func (c *Client) DeleteAssetData(ctx context.Context, v *dataproxypb.AssetData, options ...Option) error {
	// TODO with options

	_, err := c.c.DeleteAssetData(ctx, &dataproxypb.DeleteAssetDataRequest{
		AssetData: v,
	})
	return err
}

func (c *Client) GenerateSignedUrl(ctx context.Context, v *dataproxypb.AssetData, options ...Option) (string, error) {
	// TODO with options

	res, err := c.c.GenerateSignedUrl(ctx, &dataproxypb.SignedUrlRequest{
		AssetData: v,
	})
	if err != nil {
		return "", err
	}
	return res.SignedUrl, nil
}
