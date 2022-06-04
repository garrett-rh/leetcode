package prefix

func LongestCommonPrefix(strs []string) string {
	var result string

	for i := 0; true; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return result
			}
		}
		result += string(strs[0][i])
	}

	return result
}
