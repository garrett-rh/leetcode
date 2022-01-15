package goodPairs

import (
	"fmt"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	pairsTest := []struct {
		numbers []int
		pairs   int
	}{
		{numbers: []int{1, 2, 3, 1, 1, 3}, pairs: 4},
		{numbers: []int{1, 1, 1, 1}, pairs: 6},
		{numbers: []int{1, 2, 3}, pairs: 0},
	}

	for _, test := range pairsTest {
		t.Run(fmt.Sprintf("%v has %d good pairs", test.numbers, test.pairs), func(t *testing.T) {
			got := NumIdenticalPairs(test.numbers)
			if got != test.pairs {
				t.Errorf("got %v want %v", got, test.pairs)
			}
		})
	}

}
