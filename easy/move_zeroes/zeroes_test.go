package moveZeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("put zeros in back %v", test.nums), func(t *testing.T) {
			MoveZeroes(test.nums) //in place op

			if !reflect.DeepEqual(test.nums, test.want) {
				t.Errorf("got %v, want %v", test.nums, test.want)
			}
		})
	}
}
