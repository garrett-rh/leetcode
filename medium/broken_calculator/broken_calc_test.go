package brokenCalc

import (
	"fmt"
	"testing"
)

func TestBrokenCalc(t *testing.T) {
	testCases := []struct {
		startValue int
		target     int
		want       int
	}{
		{startValue: 2, target: 3, want: 2},
		{startValue: 5, target: 8, want: 2},
		{startValue: 3, target: 10, want: 3},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Started with %d, want target %d in %d ops", test.startValue, test.target, test.want), func(t *testing.T) {
			got := brokenCalc(test.startValue, test.target)

			if test.want != got {
				t.Errorf("wanted %d, got %d", test.want, got)
			}
		})
	}
}
