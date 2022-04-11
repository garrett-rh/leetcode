package baseball

import (
	"fmt"
	"testing"
)

func TestCalPoints(t *testing.T) {
	testCases := []struct {
		ops  []string
		want int
	}{
		{ops: []string{"5", "2", "C", "D", "+"}, want: 30},
		{ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, want: 27},
		{ops: []string{"1"}, want: 1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Score of %v", test.ops), func(t *testing.T) {
			got := CalPoints(test.ops)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
