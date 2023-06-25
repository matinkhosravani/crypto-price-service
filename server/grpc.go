package server

import (
	"context"
	"github.com/matinkhosravani/crypto-price-fetcher/proto"
	"github.com/matinkhosravani/crypto-price-fetcher/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	proto.UnimplementedPriceFetcherServer
	lisAddr string
	Pf      service.PriceFetcher
}

func NewGRPCServer(pf service.PriceFetcher, listenAddr string) *GRPCServer {
	return &GRPCServer{Pf: pf, lisAddr: listenAddr}
}

func (s *GRPCServer) Fetch(ctx context.Context, input *proto.PriceInput) (*proto.PriceOutput, error) {
	price, err := s.Pf.Fetch(ctx, input.Symbol)
	if err != nil {
		return nil, err
	}

	return &proto.PriceOutput{
		Symbol: input.Symbol,
		Price:  price,
	}, err
}

func (s *GRPCServer) Run() {
	ls, err := net.Listen("tcp", s.lisAddr)
	defer ls.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := grpc.NewServer()
	proto.RegisterPriceFetcherServer(srv, s)
	err = srv.Serve(ls)
	if err != nil {
		log.Fatal(err)
	}

}
