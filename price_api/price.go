package price_api

import (
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
)

const CoinMarketCapAPIType = "coinmarketcap"

type PriceAPI interface {
	Fetch() ([]types.CoinPrice, error)
}

func NewPriceAPI() PriceAPI {
	switch app.GetEnv().PriceAPI {
	case CoinMarketCapAPIType:
		return NewCoinMarketCapAPI()
	default:
		return nil
	}
}
