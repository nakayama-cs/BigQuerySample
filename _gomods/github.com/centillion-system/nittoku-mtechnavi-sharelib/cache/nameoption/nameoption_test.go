package nameoption_test

import (
	"context"
	programoptionpb "mtechnavi/programoption/protobuf"
	"mtechnavi/sharelib/cache/nameoption"
	"mtechnavi/sharelib/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestLookUpByCode(t *testing.T) {
	t.Run("LookupByCode 正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption, err := c.LookupByCode(ctx, "TenantA", "CategoryNameA", "0")
		assert.Nil(t, err)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "0", nameoption.Code)

		nameoption, err = c.LookupByCode(ctx, "TenantA", "CategoryNameA", "0")

		assert.Nil(t, err)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "0", nameoption.Code)
	})

	t.Run("LookupByCode 準正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption, err := c.LookupByCode(ctx, "TenantA", "CategoryNameC", "0")
		assert.NotNil(t, err)
		assert.Nil(t, nameoption)

		nameoption, err = c.LookupByCode(ctx, "TenantA", "CategoryNameA", "2")
		assert.NotNil(t, err)
		assert.Nil(t, nameoption)
	})
}

func TestLookUpCategory(t *testing.T) {
	t.Run("LookupCategory 正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoptions, err := c.LookupCategory(ctx, "TenantA", "CategoryNameA")
		assert.Nil(t, err)
		assert.Equal(t, 2, len(nameoptions))
		assert.Equal(t, "CategoryNameA", nameoptions[0].CategoryName)
		assert.Equal(t, "0", nameoptions[0].Code)
		assert.Equal(t, "CategoryNameA", nameoptions[1].CategoryName)
		assert.Equal(t, "1", nameoptions[1].Code)
	})

	t.Run("LookupCategory 準正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption, err := c.LookupCategory(ctx, "TenantA", "CategoryNameC")
		assert.Nil(t, err)
		assert.Nil(t, nameoption)
	})
}

func TestLookupByCodeWithDefault(t *testing.T) {
	t.Run("LookupByCodeWithDefault 正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption := c.LookupByCodeWithDefault(ctx, "TenantA", "CategoryNameA", "0", nil)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "0", nameoption.Code)

		nameoption = c.LookupByCodeWithDefault(ctx, "TenantA", "CategoryNameA", "0", nil)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "0", nameoption.Code)

		nameoption = c.LookupByCodeWithDefault(ctx, "TenantA", "CategoryNameA", "5", nil)
		assert.Nil(t, nameoption)

		nameoption = c.LookupByCodeWithDefault(ctx, "TenantA", "CategoryNameA", "5", &protobuf.NameOption{CategoryName: "Default"})
		assert.NotNil(t, nameoption)
		assert.Equal(t, "Default", nameoption.CategoryName)
	})
}

func TestLookUpBySystemName(t *testing.T) {
	t.Run("LookupBySystemName 正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption, err := c.LookupBySystemName(ctx, "TenantA", "CategoryNameA", "B00")
		assert.Nil(t, err)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "B00", nameoption.SystemName)

		nameoption, err = c.LookupBySystemName(ctx, "TenantA", "CategoryNameA", "B00")

		assert.Nil(t, err)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "B00", nameoption.SystemName)
	})

	t.Run("LookupBySystemCode 準正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption, err := c.LookupByCode(ctx, "TenantA", "CategoryNameC", "B00")
		assert.NotNil(t, err)
		assert.Nil(t, nameoption)

		nameoption, err = c.LookupByCode(ctx, "TenantA", "CategoryNameA", "B02")
		assert.NotNil(t, err)
		assert.Nil(t, nameoption)
	})
}

func TestLookupBySystemNameWithDefault(t *testing.T) {
	t.Run("LookupBySystemNameWithDefault 正常系", func(t *testing.T) {
		ctx := context.Background()

		cli := DummyClient{}

		c := &nameoption.Cache{
			CacheSize: 1024,
			Client:    &cli,
		}
		if err := c.Init(ctx); err != nil {
			panic(err)
		}

		nameoption := c.LookupBySystemNameWithDefault(ctx, "TenantA", "CategoryNameA", "B00", nil)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "B00", nameoption.SystemName)

		nameoption = c.LookupBySystemNameWithDefault(ctx, "TenantA", "CategoryNameA", "B00", nil)
		assert.Equal(t, "CategoryNameA", nameoption.CategoryName)
		assert.Equal(t, "B00", nameoption.SystemName)

		nameoption = c.LookupBySystemNameWithDefault(ctx, "TenantA", "CategoryNameA", "B05", nil)
		assert.Nil(t, nameoption)

		nameoption = c.LookupBySystemNameWithDefault(ctx, "TenantA", "CategoryNameA", "B05", &protobuf.NameOption{CategoryName: "Default"})
		assert.NotNil(t, nameoption)
		assert.Equal(t, "Default", nameoption.CategoryName)
	})
}

type DummyClient struct {
}

func (c *DummyClient) ListProgramOptions(context.Context, *programoptionpb.ListProgramOptionsRequest, ...grpc.CallOption) (*programoptionpb.ListProgramOptionsResponse, error) {
	result := &programoptionpb.ListProgramOptionsResponse{}

	result.Items = []*programoptionpb.ProgramOption{
		{
			ProgramOptionId: "POIA01",
			CategoryName:    "CategoryNameA",
			Code:            "0",
			DisplayNameLang: map[string]string{
				"ja": "JapaneseA01",
			},
			SystemName: "B00",
			System: &programoptionpb.ProgramOption_DisplayNameLang{
				DisplayNameLang: map[string]string{
					"ja": "JapaneseA01",
				},
			},
			Order: 1,
		},
		{
			ProgramOptionId: "POIA02",
			CategoryName:    "CategoryNameA",
			Code:            "1",
			DisplayNameLang: map[string]string{
				"ja": "JapaneseA02",
			},
			SystemName: "B01",
			System: &programoptionpb.ProgramOption_DisplayNameLang{
				DisplayNameLang: map[string]string{
					"ja": "JapaneseA02",
				},
			},
			Order: 2,
		},
		{
			ProgramOptionId: "POIB01",
			CategoryName:    "CategoryNameB",
			Code:            "0",
			DisplayNameLang: map[string]string{
				"ja": "JapaneseB01",
			},
			SystemName: "B00",
			System: &programoptionpb.ProgramOption_DisplayNameLang{
				DisplayNameLang: map[string]string{
					"ja": "JapaneseB01",
				},
			},
			Order: 1,
		},
	}

	result.Total = int64(len(result.Items))

	return result, nil
}
