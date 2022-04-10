package topKFrequentElements

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{nums: []int{1}, k: 1, want: []int{1}},
		{nums: []int{5, 2, 3, 5, 2, 2, 1}, k: 2, want: []int{2, 5}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Most common elements in %v", test.nums), func(t *testing.T) {
			got := TopKFrequent(test.nums, test.k)
			sort.Ints(got)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
