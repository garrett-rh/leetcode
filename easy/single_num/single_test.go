package single

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{4, 1, 2, 1, 2}, want: 4},
		{nums: []int{1}, want: 1},
		{nums: []int{2, 2, 1}, want: 1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Find the unique number in %v", test.nums), func(t *testing.T) {
			got := SingleNumber(test.nums)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
