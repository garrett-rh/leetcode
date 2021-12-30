package wordsinsentences

import "strings"

func MostWordsFound(sentences []string) int {
	var ans int
	for i := 0; i < len(sentences); i++ {
		if len(strings.Split(sentences[i], " ")) > ans {
			ans = len(strings.Split(sentences[i], " "))
		}
	}
	return ans
}
