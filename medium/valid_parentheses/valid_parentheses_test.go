package validParentheses

import (
	"fmt"
	"testing"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	testCases := []struct {
		given string
		want  string
	}{
		{given: "lee(t(c)o)de)", want: "lee(t(c)o)de"},
		{given: "a)b(c)d", want: "ab(c)d"},
		{given: "))((", want: ""},
		{given: "he(ll)o", want: "he(ll)o"},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Removal of parenthesis from %s", test.given), func(t *testing.T) {
			got := MinRemoveToMakeValid(test.given)
			if got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}

}
