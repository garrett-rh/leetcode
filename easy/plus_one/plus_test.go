package plus

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPluSOne(t *testing.T) {
	testCases := []struct {
		digits []int
		want   []int
	}{
		{digits: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		{digits: []int{9}, want: []int{1, 0}},
		{digits: []int{9, 9, 9}, want: []int{1, 0, 0, 0}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("increment by one of %v", test.digits), func(t *testing.T) {
			got := PlusOne(test.digits)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
