package types

type PriceResponse struct {
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
}

type CoinPrice struct {
	Symbol string
	Price  float64
}
