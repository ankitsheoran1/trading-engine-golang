package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBidAndAsk(t *testing.T) {
	engine := newMatchEngine()
	engine.addNewMarket(&TradingPair{base: "BTC", quote: "USD"})

	// Add for Ask 7 order
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 8, &Order{orderType: Ask, size: 4})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Ask, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 10, &Order{orderType: Ask, size: 1})

	// Add for Ask 8 order
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Bid, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Bid, size: 1})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 10, &Order{orderType: Bid, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 11, &Order{orderType: Bid, size: 3})

	processAskOrder(engine)

	// total ask order
	askCount := 0
	bidCount := 0
	for _, orderbook := range engine.orderbooks {
		for _, limit := range orderbook.ask {
			for _, order := range limit.orders {
				askCount += order.size
			}
		}

		for _, limit := range orderbook.bid {
			for _, order := range limit.orders {
				bidCount += order.size
			}
		}
	}

	assert.Equal(t, askCount, 0, "askCount is not matching")
	assert.Equal(t, bidCount, 1, "bidCount is not matching")
}
