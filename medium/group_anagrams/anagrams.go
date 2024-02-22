package anagrams

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	sorter := make(map[string][]string)

	for _, word := range strs {
		wordSlice := strings.Split(word, "")
		slices.Sort(wordSlice)
		sorted_string := strings.Join(wordSlice, "")
		sorter[sorted_string] = append(sorter[sorted_string], word)
	}

	ans := [][]string{}
	for _, v := range sorter {
		ans = append(ans, v)
	}

	return ans
}
