package remove

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		nums   []int
		result []int
		val    int
		want   int
	}{
		{nums: []int{3, 2, 2, 3}, result: []int{2, 2}, val: 3, want: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, result: []int{0, 0, 1, 3, 4}, val: 2, want: 5},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("remove %d from %v", test.val, test.nums), func(t *testing.T) {
			got := RemoveElement(test.nums, test.val)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}

			sort.Ints(test.nums)
			if !reflect.DeepEqual(test.nums[:got], test.result) {
				t.Errorf("%d not removed from %v to match %v", test.val, test.nums, test.result)
			}
		})
	}
}
