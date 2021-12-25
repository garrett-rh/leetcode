package richest_customer

import (
	"sort"
)

func MaximumWealth(accounts [][]int) int {
	var wealth []int
	for _, account := range accounts {
		var res int
		for i := 0; i < len(account); i++ {
			res += account[i]
		}
		wealth = append(wealth, res)
	}

	sort.Ints(wealth)
	return wealth[len(wealth)-1]
}
