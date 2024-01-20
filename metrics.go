package main

import (
	"context"
	"fmt"
)

type MetricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &MetricService{
		next: next,
	}
}

func (ms *MetricService) FetchPrice(ctx context.Context, symbol string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus")
	return ms.next.FetchPrice(ctx, symbol)
}
