package issubsequence

func isSubsequence(s string, t string) bool {
	// exit immediately if length is 0
	if len(s) == 0 {
		return true
	}
	// use i to increment through the longer (t)
	// use j to track the place if correct
	var j int
	for i := 0; i < len(t); i++ {
		// if s[j] == t[i] then we can increment to the next character in s. otherwise, stay at the same
		if s[j] == t[i] {
			j++
		}
		// return true if all characters of s have been accounted for
		if j == len(s) {
			return true
		}
	}

	// if it hasn't returned by now, s is not a subsequence of t
	return false
}
