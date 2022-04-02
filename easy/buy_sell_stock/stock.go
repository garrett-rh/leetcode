package gmeToTheMoon

func MaxProfit(prices []int) int {

	var (
		lowest  = prices[0]
		highest = prices[0]
		profit  int
	)

	for _, price := range prices {
		if price > highest {
			highest = price
		}
		if price < lowest {
			lowest = price
			highest = lowest
		}

		if highest-lowest > profit {
			profit = highest - lowest
		}
	}

	return profit
}
