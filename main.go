package main

import (
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/server"
	"github.com/matinkhosravani/crypto-price-fetcher/service"
)

func main() {
	app.Boot()
	gRPCServer := server.NewGRPCServer(service.NewLogging(service.NewPriceFetcher()), app.GetEnv().GRPCListenAddr)
	go gRPCServer.Run()
	jsonServer := server.NewJsonServer(app.GetEnv().JSONListenAddr)
	jsonServer.Run()
}
