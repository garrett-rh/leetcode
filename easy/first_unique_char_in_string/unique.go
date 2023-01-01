package firstUnique

func FirstUniqChar(s string) int {
	m := make([]int, 26)

	for _, char := range s {
		m[int(char-'a')]++
	}

	for idx, char := range s {
		if m[int(char-'a')] == 1 {
			return idx
		}
	}

	return -1
}
