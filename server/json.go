package server

import (
	"github.com/matinkhosravani/crypto-price-fetcher/handler"
	"github.com/matinkhosravani/crypto-price-fetcher/service"
	"log"
	"net/http"
)

type JsonServer struct {
	listenAddr string
}

func NewJsonServer(listenAddr string) *JsonServer {
	return &JsonServer{listenAddr: listenAddr}
}

func (s JsonServer) Run() {
	h := handler.PriceHandler{
		Pf: service.NewLogging(service.NewPriceFetcher()),
	}

	http.HandleFunc("/price", h.GetPrice)
	err := http.ListenAndServe(s.listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
