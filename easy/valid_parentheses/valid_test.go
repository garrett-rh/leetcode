package validParentheses

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(]", want: false},
		{s: "{[]}", want: true},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Valid brackets/parentheses in %s", test.s), func(t *testing.T) {

			got := IsValid(test.s)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
