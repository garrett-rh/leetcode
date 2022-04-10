package removeDuplicates

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 2}, want: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Remove Duplicates in %v", test.nums), func(t *testing.T) {
			got := RemoveDuplicates(test.nums)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
