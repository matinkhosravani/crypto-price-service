package price_api

import (
	"encoding/json"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
	"io"
	"net/http"
	"net/url"
)

type coinMarketCapAPI struct{}

func NewCoinMarketCapAPI() *coinMarketCapAPI {
	return &coinMarketCapAPI{}
}

type coinMarketCapCryptos struct {
	Data struct {
		CryptoCurrencyList []struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			Slug   string `json:"slug"`
			Quotes []struct {
				Price float64 `json:"price"`
			} `json:"quotes"`
		} `json:"cryptoCurrencyList"`
	} `json:"data"`
}

func (cmc coinMarketCapAPI) Fetch() ([]types.CoinPrice, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")
	q.Add("sortBy", "market_cap")
	q.Add("sortType", "desc")
	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := io.ReadAll(resp.Body)
	var listing coinMarketCapCryptos
	err = json.Unmarshal(respBody, &listing)
	if err != nil {
		return nil, err
	}
	var coinPrizes []types.CoinPrice
	for _, coin := range listing.Data.CryptoCurrencyList {
		coinPrizes = append(coinPrizes, types.CoinPrice{
			Symbol: coin.Symbol,
			Price:  coin.Quotes[0].Price,
		})
	}

	return coinPrizes, nil
}
