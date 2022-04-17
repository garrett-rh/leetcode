package rotate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
		{nums: []int{1}, k: 1, want: []int{1}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Rotate %v at index %d", test.nums, test.k), func(t *testing.T) {
			Rotate(test.nums, test.k)
			if !reflect.DeepEqual(test.nums, test.want) {
				t.Errorf("got : %v, want : %v", test.nums, test.want)
			}
		})
	}
}
