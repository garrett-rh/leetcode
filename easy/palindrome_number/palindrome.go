package palindrome

import "strconv"

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	for i, j := len(xStr)-1, 0; i > 0; i, j = i-1, j+1 {
		if i == j {
			break
		}
		if xStr[i] != xStr[j] {
			return false
		}
	}

	return true
}
