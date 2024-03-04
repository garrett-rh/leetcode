package mergealternately

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		word1 string
		word2 string
		want  string
	}{
		{word1: "abc", word2: "pqr", want: "apbqcr"},
		{word1: "ab", word2: "pqrs", want: "apbqrs"},
		{word1: "abcd", word2: "pq", want: "apbqcd"},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Merging %s and %s", test.word1, test.word2), func(t *testing.T) {

			if got := mergeAlternately(test.word1, test.word2); got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}
