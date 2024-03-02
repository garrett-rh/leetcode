package sortedsquares

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
		{nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
	}

	for _, test := range testCases {

		t.Run(fmt.Sprintf("Sorted square roots of %v", test.nums), func(t *testing.T) {
			got := sortedSquares(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
