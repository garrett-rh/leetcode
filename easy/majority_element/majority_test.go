package majority

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{3, 2, 3}, want: 3},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("majority num in array %v", test.nums), func(t *testing.T) {
			got := MajorityElement(test.nums)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
