package commondivisorstrings

import (
	"fmt"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		want string
	}{
		{str1: "ABCABC", str2: "ABC", want: "ABC"},
		{str1: "ABABAB", str2: "ABAB", want: "AB"},
		{str1: "LEFT", str2: "CODE", want: ""},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Greatest Common Divisor of %s and %s", test.str1, test.str2), func(t *testing.T) {

			if got := gcdOfStrings(test.str1, test.str2); got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}
