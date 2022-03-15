package jewelsAndStones

import (
	"strings"
)

func NumJewelsInStones(jewels string, stones string) int {

	result := 0
	for idx := range stones {
		if strings.Contains(jewels, string(stones[idx])) {
			result++
		}
	}

	return result
}
