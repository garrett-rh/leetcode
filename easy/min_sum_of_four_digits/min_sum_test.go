package minSum

import (
	"fmt"
	"testing"
)

func TestMinimumSum(t *testing.T) {
	minSum := []struct {
		number int
		sum    int
	}{
		{number: 2392, sum: 52},
		{number: 4009, sum: 13},
	}

	for _, test := range minSum {
		t.Run(fmt.Sprintf("%d has a minimum sum of %d", test.number, test.sum), func(t *testing.T) {
			got := MinimumSum(test.number)
			if got != test.sum {
				t.Errorf("got %d, want %d", got, test.sum)
			}
		})
	}

}
