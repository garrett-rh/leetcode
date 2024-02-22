package anagrams

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRestoreArray(t *testing.T) {
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{strs: []string{""}, want: [][]string{{""}}},
		{strs: []string{"a"}, want: [][]string{{"a"}}},
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("sorting anagrams of %v", test.strs), func(t *testing.T) {
			got := groupAnagrams(test.strs)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
