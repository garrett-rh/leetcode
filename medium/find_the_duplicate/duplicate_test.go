package duplicates

import (
	"fmt"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 4, 2, 2}, want: 2},
		{nums: []int{3, 1, 3, 4, 2}, want: 3},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Array %v has a duplicate value of %d", test.nums, test.want), func(t *testing.T) {
			got := FindDuplicates(test.nums)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
