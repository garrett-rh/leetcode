package stonks

import (
	"fmt"
	"testing"
)

func TestMaxProfits(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{prices: []int{1, 2, 3, 4, 5}, want: 4},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Buy & Sell stock: %v", test.prices), func(t *testing.T) {
			got := MaxProfits(test.prices)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
