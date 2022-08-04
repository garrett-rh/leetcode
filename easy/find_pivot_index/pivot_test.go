package pivot

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{nums: []int{1, 2, 3}, want: -1},
		{nums: []int{2, 1, -1}, want: 0},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Pivot index of %v", test.nums), func(t *testing.T) {
			got := pivotIndex(test.nums) //in place op

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
