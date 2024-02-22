package adjacent

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRestoreArray(t *testing.T) {
	testCases := []struct {
		pairs [][]int
		want  []int
	}{
		{pairs: [][]int{{2, 1}, {3, 4}, {3, 2}}, want: []int{1, 2, 3, 4}},
		{pairs: [][]int{{4, -2}, {1, 4}, {-3, 1}}, want: []int{-2, 4, 1, -3}},
		{pairs: [][]int{{10000, -10000}}, want: []int{10000, -10000}}}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("restoring %v array to adjacent pairs", test.pairs), func(t *testing.T) {
			got := restoreArray(test.pairs)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("got %v, want %v", got, test.want)
			}

		})
	}
}
