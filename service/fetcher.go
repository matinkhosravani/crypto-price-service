package service

import (
	"context"
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/cache"
	"github.com/matinkhosravani/crypto-price-fetcher/price_api"
	"time"
)

type PriceFetcher interface {
	Fetch(ctx context.Context, symbol string) (float64, error)
}

type priceFetcher struct {
	c   cache.Cache
	api price_api.PriceAPI
}

func NewPriceFetcher() *priceFetcher {
	return &priceFetcher{
		c:   cache.NewCache(),
		api: price_api.NewPriceAPI(),
	}
}

func (p priceFetcher) Fetch(ctx context.Context, symbol string) (float64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	price, exists, err := p.c.GetPriceBySymbol(ctx, symbol)
	if err != nil {
		return -1, err
	}

	if !exists {
		coins, err := p.api.Fetch()
		if err != nil {
			return -1, err
		}

		exp := time.Duration(app.GetEnv().PriceExpirationTime) * time.Minute
		err = p.c.SetCoins(ctx, coins, exp)
		if err != nil {
			return -1, err
		}

		price, _, err = p.c.GetPriceBySymbol(ctx, symbol)
		if err != nil {
			return -1, err
		}
	}

	return price, err
}
