package romanToInt

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{s: "III", want: 3},
		{s: "LVIII", want: 58},
		{s: "MCMXCIV", want: 1994},
		{s: "IV", want: 4},
		{s: "XIV", want: 14},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Roman numeral %s to %d", test.s, test.want), func(t *testing.T) {
			got := RomanToInt(test.s)

			if test.want != got {
				t.Errorf("wanted %d, got %d", test.want, got)
			}
		})
	}
}
