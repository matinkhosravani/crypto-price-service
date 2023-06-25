package main

import (
	"github.com/matinkhosravani/crypto-price-fetcher/app"
	"github.com/matinkhosravani/crypto-price-fetcher/server"
)

func main() {
	app.Boot()

	jsonServer := server.NewJsonServer(app.GetEnv().JSONListenAddr)
	jsonServer.Run()
}
