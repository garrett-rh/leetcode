package strstr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{haystack: "hello", needle: "ll", want: 2},
		{haystack: "aaaaa", needle: "bba", want: -1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("needle %s in haystack %s", test.needle, test.haystack), func(t *testing.T) {
			got := StrStr(test.haystack, test.needle)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
