package firstUnique

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{s: "leetcode", want: 0},
		{s: "loveleetcode", want: 2},
		{s: "aabb", want: -1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("first unique char in %s", test.s), func(t *testing.T) {
			got := FirstUniqChar(test.s)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
