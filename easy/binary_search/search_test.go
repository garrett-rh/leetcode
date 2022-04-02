package binarySearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Binary search of %v, target %d", test.nums, test.target), func(t *testing.T) {
			got := Search(test.nums, test.target)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
