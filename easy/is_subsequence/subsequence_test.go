package issubsequence

import (
	"fmt"
	"testing"
)

func TestIsSubSequence(t *testing.T) {
	testCases := []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abc", t: "ahbgdc", want: true},
		{s: "axc", t: "ahbgdc", want: false},
		{s: "", t: "ahbgdc", want: true},
		{s: "b", t: "abc", want: true},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test if %s is a subsequence of %s", test.s, test.t), func(t *testing.T) {

			if got := isSubsequence(test.s, test.t); got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
