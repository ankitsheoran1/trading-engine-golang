package main

const (
	Bid = iota + 1
	Ask
)

type OrderBook struct {
	ask map[int]Limit
	bid map[int]Limit
}

func (o *OrderBook) addOrder(price int, symbol string, order *Order) {
	if order.orderType == Bid {
		if limit, ok := o.bid[price]; ok {
			limit.orders = append(limit.orders, *order)
			o.bid[price] = limit
		} else {
			limit := Limit{
				price:  price,
				symbol: symbol,
				orders: make([]Order, 0),
			}
			limit.addOrder(order)
			o.bid[price] = limit

		}
	} else {
		if limit, ok := o.ask[price]; ok {
			limit.addOrder(order)
			o.ask[price] = limit
		} else {
			limit := Limit{
				price:  price,
				symbol: symbol,
				orders: make([]Order, 0),
			}
			limit.addOrder(order)
			o.ask[price] = limit

		}
	}
}

// TODO Store asks and bid in sorted order  of time and get direct appropriate limit rather thn traverse
func (o *OrderBook) fillOrder(order *Order, price int) (Order, error) {
	if order.orderType == Bid {
		for k, v := range o.ask {
			// fmt.Println("serving bid", k, v)
			if k >= price {
				order, err := v.fillOrder(order)
				if order.isFilled() {
					return order, err
				}
			}
		}

		return *order, nil
	} else {
		for k, v := range o.bid {
			// fmt.Println("serving ask", k, v)
			if k >= price {
				order, err := v.fillOrder(order)
				if order.isFilled() {
					return order, err
				}
			}

		}

		return *order, nil

	}
}

type Limit struct {
	price  int
	symbol string
	orders []Order
}

type Order struct {
	orderType int
	size      int
}

func (o *Order) isFilled() bool {
	return o.size == 0
}

func (l *Limit) volume() int {
	count := 0
	for _, order := range l.orders {
		count += order.size
	}
	return count
}

func (l *Limit) addOrder(order *Order) {
	l.orders = append(l.orders, *order)
}

func (l *Limit) fillOrder(matchOrder *Order) (Order, error) {

	for i, order := range l.orders {
		/// fmt.Println("orders ", l.orders)
		if order.size >= matchOrder.size {
			// fmt.Println("order size have sufficient size ", *matchOrder)
			order.size = order.size - matchOrder.size
			matchOrder.size = 0
			l.orders[i] = order
		} else {
			// fmt.Println("order size have not sufficient size ", *matchOrder)
			matchOrder.size = matchOrder.size - order.size
			order.size = 0
			l.orders[i] = order
		}
		if matchOrder.isFilled() {
			break
		}

	}
	return *matchOrder, nil
}
