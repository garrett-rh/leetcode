package searchInsert

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, want: 4},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Return index of where %d is or would be in %v", test.target, test.nums), func(t *testing.T) {
			got := SearchInsert(test.nums, test.target)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
