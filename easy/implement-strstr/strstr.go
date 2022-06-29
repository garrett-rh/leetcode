package strstr

import "strings"

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	return strings.Index(haystack, needle)
}
