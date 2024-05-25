package main

import (
	"fmt"
	"time"
)

func processAskOrder(engine *MatchEngine) {
	fmt.Println("Before processing ordersbook ", engine.orderbooks)
	for _, orderbook := range engine.orderbooks {
		for price, askLimit := range orderbook.ask {
			for i, order := range askLimit.orders {
				ord, err := engine.placeLimitOrder(TradingPair{"BTC", "USD"}, price, &order)
				if err == nil {
					askLimit.orders[i] = ord
				}
			}
		}
	}

	fmt.Println("After processing ordersbook ", engine.orderbooks)

}

func ProcessOrders(engine *MatchEngine) {
	for {
		<-time.After(2 * 1000 * 1000 * 1000)
		// Get all ask order
		processAskOrder(engine)

	}
}

func main() {
	engine := newMatchEngine()
	engine.addNewMarket(&TradingPair{base: "BTC", quote: "USD"})
	// engine.Debug()
	//limit1 := Limit{price: 10, symbol: "BTC", orders: make([]Order, 0)}
	//limit2 := Limit{price: 9, symbol: "BTC", orders: make([]Order, 0)}
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 8, &Order{orderType: Ask, size: 4})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Ask, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 10, &Order{orderType: Ask, size: 1})

	//engine.Debug()
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Bid, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Bid, size: 1})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 10, &Order{orderType: Bid, size: 2})
	engine.AddOrder(TradingPair{base: "BTC", quote: "USD"}, 11, &Order{orderType: Bid, size: 3})
	// fmt.Println("=error is == ", err)
	//engine.AddOrder(&Order{orderType: Ask, size: 2})
	// engine.Debug()

	processAskOrder(engine)

	//engine.Debug()

	//engine.placeLimitOrder(TradingPair{base: "BTC", quote: "USD"}, 7, &Order{orderType: Bid, size: 3})
	//engine.placeLimitOrder(TradingPair{base: "BTC", quote: "USD"}, 9, &Order{orderType: Ask, size: 2})
	//engine.Debug()

}
