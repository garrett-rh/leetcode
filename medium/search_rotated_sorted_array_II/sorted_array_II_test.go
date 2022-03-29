package rotatedAndSortedSearchII

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   bool
	}{
		{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0, want: true},
		{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3, want: false},
		{nums: []int{1, 0, 1, 1, 1}, target: 0, want: true},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Array %v, target %d", test.nums, test.target), func(t *testing.T) {
			got := Search(test.nums, test.target)

			if got != test.want {
				t.Errorf("got %t, want %t", got, test.want)
			}

		})
	}
}
