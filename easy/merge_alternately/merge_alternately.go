package mergealternately

import "fmt"

func mergeAlternately(word1 string, word2 string) string {

	// setting idx outside the loop to use later on in the final return
	var (
		idx int
		buf string
	)

	for idx = range word1 {
		// if idx ever equals the len of word2, return buffer + rest of word 1
		if idx == len(word2) {
			return buf + word1[idx:]
		}
		buf += fmt.Sprintf("%c%c", word1[idx], word2[idx])
	}

	return buf + word2[idx+1:]
}
