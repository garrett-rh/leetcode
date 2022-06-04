package prefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		strs []string
		want string
	}{
		{strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{strs: []string{"dog", "racecar", "car"}, want: ""},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("find the longest prefix in %v", test.strs), func(t *testing.T) {
			got := LongestCommonPrefix(test.strs)

			if got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}
