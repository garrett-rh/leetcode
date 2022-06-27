package deciBinary

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		n    string
		want int
	}{
		{n: "32", want: 3},
		{n: "82734", want: 8},
		{n: "27346209830709182346", want: 9},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("minimum deci binary number of %s", test.n), func(t *testing.T) {
			got := MinPartitions(test.n)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}

		})
	}
}
