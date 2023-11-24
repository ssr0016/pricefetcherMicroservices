package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher
type priceFetcher struct{}

// business logic
func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic the HTTP roundtrip
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return 0, fmt.Errorf("unknown ticker %s", ticker)
	}

	return price, nil
}
