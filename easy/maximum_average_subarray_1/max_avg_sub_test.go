package maximumaveragesubarray1

import (
	"fmt"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want float64
	}{
		{nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75},
		{nums: []int{5}, k: 1, want: 5.0000},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Find max avg subarray of %v in a contiguous array of %d", test.nums, test.k), func(t *testing.T) {

			if got := findMaxAverage(test.nums, test.k); got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
