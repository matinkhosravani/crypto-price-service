package redis

import (
	"context"
	"fmt"
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Cache struct {
	client *redis.Client
}

func NewRedisCache() (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", app.GetEnv().RedisHost, app.GetEnv().RedisPort),
		Password: app.GetEnv().RedisPassword, // no password set
		DB:       app.GetEnv().RedisDatabase, // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return &Cache{
		client: rdb,
	}, nil
}

func (c Cache) GetPriceBySymbol(ctx context.Context, symbol string) (float64, bool, error) {
	key := c.client.Exists(ctx, coinsKey())
	if key.Val() == 0 {
		return -1, false, nil
	}

	res := c.client.HGet(ctx, coinsKey(), symbol)
	if res.Err() == redis.Nil {
		return -1, false, fmt.Errorf("no such coin %s", symbol)
	} else if res.Err() != nil {
		return -1, false, res.Err()
	}

	price, err := strconv.ParseFloat(res.Val(), 64)

	return price, true, err
}

func (c Cache) SetCoins(ctx context.Context, prices []types.CoinPrice, exp time.Duration) error {
	m := make(map[string]interface{})

	for _, coin := range prices {
		m[coin.Symbol] = coin.Price
	}
	res := c.client.HSet(ctx, coinsKey(), m)
	c.client.Expire(ctx, coinsKey(), exp)

	return res.Err()
}

func coinsKey() string {
	return "coins"
}
