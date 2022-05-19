package palindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		num  int
		want bool
	}{
		{num: 121, want: true},
		{num: -121, want: false},
		{num: 10, want: false},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("number %d is the same forwards & backwards", test.num), func(t *testing.T) {
			got := IsPalindrome(test.num) //in place op

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
