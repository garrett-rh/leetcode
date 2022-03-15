package jewelsAndStones

import (
	"fmt"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewels string
		stones string
		want   int
	}{
		{jewels: "aA", stones: "aAAbbbb", want: 3},
		{jewels: "z", stones: "ZZ", want: 0},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("jewels: %s, stones: %s", test.jewels, test.stones), func(t *testing.T) {
			got := NumJewelsInStones(test.jewels, test.stones)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
