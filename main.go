package main

import "fmt"

type trade struct {
	orderId int
	stock   string
	buy     bool
	priee   float64
	amount  int
	time    string
}

func main() {
	buy := make([]trade, 0)
	sell := make([]trade, 0)
	var input []trade

	input = append(input, trade{1, false, 240.12, 100, ""})
	input = append(input, trade{2, false, 237.45, 90, ""})
	input = append(input, trade{3, true, 238.10, 110, ""})
	input = append(input, trade{4, true, 237.80, 10, ""})
	input = append(input, trade{5, true, 237.80, 40, ""})
	input = append(input, trade{6, false, 236.00, 50, ""})

	for _, v := range input {
		if v.buy {
			sellOrdes := getSellOrderIdx(sell, v)

			for _, i := range sellOrdes {
				if sell[i].amount <= v.amount {
					fmt.Println(fmt.Sprintf("#%d %f %d %d", v.orderId, sell[i].priee, sell[i].amount, sell[i].orderId))
					v.amount = v.amount - sell[i].amount
					sell[i].amount = 0
				} else {
					fmt.Println(fmt.Sprintf("#%d %f %d %d", v.orderId, sell[i].priee, sell[i].amount, sell[i].orderId))
					sell[i].amount = sell[i].amount - v.amount
					v.amount = 0
				}
			}
			if v.amount > 0 {
				buy = append(buy, v)
			}
		} else {
			buyOrders := getBuyOrderIdx(buy, v)
			for _, i := range buyOrders {
				if buy[i].amount <= v.amount {
					fmt.Println(fmt.Sprintf("#%d %f %d %d", buy[i].orderId, v.priee, buy[i].amount, v.orderId))
					v.amount = v.amount - buy[i].amount
					buy[i].amount = 0
				} else {
					fmt.Println(fmt.Sprintf("#%d %f %d %d", buy[i].orderId, v.priee, buy[i].amount, v.orderId))
					buy[i].amount = buy[i].amount - v.amount
					v.amount = 0
				}
			}
			if v.amount > 0 {
				sell = append(sell, v)
			}
		}
	}
}

func getSellOrderIdx(sell []trade, buy trade) []int {
	rem := buy.amount
	var orders []int
	for i, v := range sell {
		if v.priee <= buy.priee && v.amount > 0 {
			orders = append(orders, i)
			rem = rem - v.amount
			if rem <= 0 {
				return orders
			}
		}
	}
	return orders
}

func getBuyOrderIdx(buy []trade, sell trade) []int {
	rem := sell.amount
	var orders []int
	for i, v := range buy {
		if v.priee >= sell.priee && v.amount > 0 {
			orders = append(orders, i)
			rem = rem - v.amount
			if rem <= 0 {
				return orders
			}
		}
	}
	return orders
}
