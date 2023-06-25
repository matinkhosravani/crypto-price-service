package handler

import (
	"context"
	"github.com/matinkhosravani/crypto-price-fetcher/service"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
	"github.com/matinkhosravani/crypto-price-fetcher/util"
	"net/http"
	"strings"
)

type PriceHandler struct {
	Pf service.PriceFetcher
}

func (h PriceHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	symbol := r.URL.Query().Get("symbol")
	symbol = strings.ToUpper(symbol)

	price, err := h.Pf.Fetch(ctx, symbol)

	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Symbol: symbol,
	}

	util.WriteJSON(w, http.StatusOK, &priceResp)
}
