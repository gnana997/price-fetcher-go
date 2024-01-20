package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (pf *priceFetcher) FetchPrice(ctx context.Context, symbol string) (float64, error) {
	return pf.FetchPriceMock(ctx, symbol)
}

var priceMocks = map[string]float64{
	"BTC":  20_000.0,
	"ETH":  10_000.0,
	"LINK": 50_000.0,
}

func (pf *priceFetcher) FetchPriceMock(ctx context.Context, symbol string) (float64, error) {
	price, ok := priceMocks[symbol]
	if !ok {
		return price, fmt.Errorf("currency %s not supported", symbol)
	}
	return priceMocks[symbol], nil
}
