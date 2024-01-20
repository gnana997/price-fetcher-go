package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type logginService struct {
	next PriceFetcher
}

func NewLogginServe(next PriceFetcher) PriceFetcher {
	return &logginService{
		next: next,
	}
}

func (s *logginService) FetchPrice(ctx context.Context, symbol string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(
			logrus.Fields{
				"requestId": ctx.Value("requestId"),
				"took":      time.Since(begin),
				"err":       err,
				"price":     price,
			},
		).Info("fetched price")
	}(time.Now())
	return s.next.FetchPrice(ctx, symbol)
}
