package gmeToTheMoon

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Best day to buy & sell: %v", test.prices), func(t *testing.T) {
			got := MaxProfit(test.prices)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
