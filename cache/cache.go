package cache

import (
	"context"
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/cache/redis"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
	"log"
	"time"
)

const RedisCacheType = "redis"

type Cache interface {
	GetPriceBySymbol(ctx context.Context, symbol string) (float64, bool, error)
	SetCoins(ctx context.Context, coins []types.CoinPrice, exp time.Duration) error
}

func NewCache() Cache {
	switch app.GetEnv().CacheType {
	case RedisCacheType:
		c, err := redis.NewRedisCache()
		if err != nil {
			log.Fatal(err.Error())
		}
		return c
	default:
		return nil
	}
}
