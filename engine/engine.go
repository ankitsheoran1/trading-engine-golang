package main

import (
	"fmt"
)

// BTCUSD
// BTC => BASE
// USD => QUOTE
type TradingPair struct {
	base  string
	quote string
}

type MatchEngine struct {
	orderbooks map[string]OrderBook
}

func (m *MatchEngine) addNewMarket(market *TradingPair) {
	key := market.base + market.quote
	m.orderbooks[key] = OrderBook{
		ask: make(map[int]Limit),
		bid: make(map[int]Limit),
	}
}

func newMatchEngine() *MatchEngine {
	return &MatchEngine{
		orderbooks: make(map[string]OrderBook),
	}
}

// TODO run a ticker if match not succeed
func (m *MatchEngine) placeLimitOrder(pair TradingPair, price int, order *Order) (Order, error) {
	key := pair.base + pair.quote
	if orderbook, ok := m.orderbooks[key]; ok {
		return orderbook.fillOrder(order, price)

	} else {
		return *order, fmt.Errorf("Market is not present")
	}
	return *order, fmt.Errorf("Market is not present")
}

func (m *MatchEngine) Debug() {
	for _, v := range m.orderbooks {

		for _, limit := range v.ask {
			fmt.Println("ask ordersbook ", v)
			for _, ord := range limit.orders {
				fmt.Println("ask orders are ", ord)
			}
		}
	}

	for _, v := range m.orderbooks {
		for _, limit := range v.bid {
			fmt.Println("bid ordersbook ", v)
			for _, ord := range limit.orders {
				fmt.Println("bid orders are ", ord)
			}
		}
	}
}

func (m *MatchEngine) AddOrder(pair TradingPair, price int, order *Order) error {
	key := pair.base + pair.quote
	if orderbook, ok := m.orderbooks[key]; ok {
		orderbook.addOrder(price, pair.base, order)
	} else {
		return fmt.Errorf("Market is not present")
	}
	return nil
}
