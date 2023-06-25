package client

import (
	"context"
	"github.com/matinkhosravani/crypto-price-fetcher/proto"
	"github.com/matinkhosravani/crypto-price-fetcher/types"
	"google.golang.org/grpc"
)

type Client struct {
	dialAddr string
}

func (c *Client) Fetch(ctx context.Context, symbol string) (*types.PriceResponse, error) {
	conn, err := grpc.Dial(c.dialAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cl := proto.NewPriceFetcherClient(conn)
	input := &proto.PriceInput{Symbol: symbol}

	resp, err := cl.Fetch(ctx, input)
	if err != nil {
		return nil, err
	}

	return &types.PriceResponse{
		Price:  resp.Price,
		Symbol: symbol,
	}, nil
}
