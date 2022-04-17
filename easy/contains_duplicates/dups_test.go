package containsDuplicates

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 2, 3, 1}, want: true},
		{nums: []int{1, 2, 3, 4}, want: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Duplicate detection in %v", test.nums), func(t *testing.T) {

			if got := ContainsDuplicate(test.nums); got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
