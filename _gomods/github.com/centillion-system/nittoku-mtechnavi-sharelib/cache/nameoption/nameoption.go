package nameoption

import (
	"context"
	"fmt"
	programoptionpb "mtechnavi/programoption/protobuf"
	pb "mtechnavi/sharelib/protobuf"
	"os"
	"time"

	"sync"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	lru "github.com/tidwall/tinylru"
	"google.golang.org/grpc"
)

// サービスごとの設定ではなく、システム基盤側の設定のため、環境変数を直接読む

var (
	// NameOptionキャッシュのTTL
	CacheTTL = os.Getenv("MTECHNAVI_NAMEOPTION_CACHE_TTL")

	// for testing
	TimeNow = time.Now
)

type Cache struct {
	Client interface {
		ListProgramOptions(context.Context, *programoptionpb.ListProgramOptionsRequest, ...grpc.CallOption) (*programoptionpb.ListProgramOptionsResponse, error)
	}

	CacheSize int

	ttl time.Duration

	// tenantId => *tenantCache
	data *lru.LRU
}

type tenantCache struct {
	mu sync.RWMutex

	expiresAt time.Time

	nameOptions []*pb.NameOption
}

func (c *Cache) Init(ctx context.Context) error {
	size := c.CacheSize
	if size == 0 {
		size = 1024
	}

	// TTL未指定時は300sec、指定されてint64パースできない場合は即終了
	c.ttl = time.Duration(300 * time.Second)
	if CacheTTL != "" {
		if v, err := time.ParseDuration(CacheTTL); err != nil {
			panic(err)
		} else {
			c.ttl = v
		}
	}

	c.data = &lru.LRU{}
	c.data.Resize(size)
	return nil
}

func (c *Cache) getTenantCache(ctx context.Context, tenantID string) (*tenantCache, error) {
	for {
		val, ok := c.data.Get(tenantID)
		if ok {
			if c, ok := val.(*tenantCache); ok {
				if c.expiresAt.After(TimeNow()) {
					return c, nil
				}
			}
		}
		if err := c.Reload(ctx, tenantID); err != nil {
			return nil, err
		}
	}
}

func (c *Cache) LookupBySystemName(ctx context.Context, tenantID, categoryName, systemName string) (*pb.NameOption, error) {
	tc, err := c.getTenantCache(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	tc.mu.RLock()
	defer tc.mu.RUnlock()

	for _, v := range tc.nameOptions {
		if v.CategoryName != categoryName {
			continue
		}
		if v.SystemName != systemName {
			continue
		}
		return v, nil
	}
	return nil, fmt.Errorf("Not Found CategoryName:%s,SystemName:%s", categoryName, systemName)
}

func (c *Cache) LookupByCode(ctx context.Context, tenantID, categoryName, code string) (*pb.NameOption, error) {
	tc, err := c.getTenantCache(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	tc.mu.RLock()
	defer tc.mu.RUnlock()

	for _, v := range tc.nameOptions {
		if v.CategoryName != categoryName {
			continue
		}
		if v.Code != code {
			continue
		}
		return v, nil
	}
	return nil, fmt.Errorf("Not Found CategoryName:%s,Code:%s", categoryName, code)
}

// 取得できなかった場合はDefaultを返却
func (c *Cache) LookupBySystemNameWithDefault(ctx context.Context, tenantID, categoryName, systemName string, defaultOption *pb.NameOption) *pb.NameOption {
	logger := ctxlogrus.Extract(ctx)

	if op, err := c.LookupBySystemName(ctx, tenantID, categoryName, systemName); err != nil {
		logger.WithError(err).Infof("NameOption is not found, use default %v: categoryName=%q, systemName=%q", defaultOption, categoryName, systemName)
		return defaultOption
	} else {
		return op
	}
}

// 取得できなかった場合はDefaultを返却
func (c *Cache) LookupByCodeWithDefault(ctx context.Context, tenantID, categoryName, code string, defaultOption *pb.NameOption) *pb.NameOption {
	logger := ctxlogrus.Extract(ctx)

	if op, err := c.LookupByCode(ctx, tenantID, categoryName, code); err != nil {
		logger.WithError(err).Infof("NameOption is not found, use default %v: categoryName=%q, code=%q", defaultOption, categoryName, code)
		return defaultOption
	} else {
		return op
	}
}

func (c *Cache) LookupCategory(ctx context.Context, tenantID, categoryName string) ([]*pb.NameOption, error) {
	tc, err := c.getTenantCache(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	tc.mu.RLock()
	defer tc.mu.RUnlock()

	var out []*pb.NameOption
	for _, v := range tc.nameOptions {
		if v.CategoryName != categoryName {
			continue
		}
		out = append(out, v)
	}
	return out, nil
}

func (c *Cache) Reload(ctx context.Context, tenantID string) error {
	tc := &tenantCache{}

	// ProgramOptionの全てのデータを取得し、キャッシング
	response, err := c.Client.ListProgramOptions(ctx, &programoptionpb.ListProgramOptionsRequest{})
	if err != nil {
		return err
	}

	nameoptions := []*pb.NameOption{}
	for _, v := range response.Items {
		nameoptions = append(nameoptions, &pb.NameOption{
			CategoryName:    v.CategoryName,
			Code:            v.Code,
			DisplayNameLang: v.DisplayNameLang,
			SystemName:      v.SystemName,
		})
	}

	tc.nameOptions = nameoptions
	tc.expiresAt = TimeNow().Add(c.ttl)
	c.data.Set(tenantID, tc)

	return nil
}
