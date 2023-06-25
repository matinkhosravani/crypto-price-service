package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Logging struct {
	next PriceFetcher
}

func NewLogging(next PriceFetcher) *Logging {
	return &Logging{next: next}
}

func (l Logging) Fetch(ctx context.Context, symbol string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"err":   err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now())

	return l.next.Fetch(ctx, symbol)
}
