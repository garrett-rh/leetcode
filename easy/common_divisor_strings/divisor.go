package commondivisorstrings

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:func(len1 int, len2 int) int {

		for len2 != 0 {
			len1, len2 = len2, len1%len2
		}

		return len1
	}(len(str1), len(str2))]
}
