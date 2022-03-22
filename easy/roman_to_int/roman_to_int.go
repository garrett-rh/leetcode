package romanToInt

func RomanToInt(s string) int {

	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && romanNumerals[s[i]] < romanNumerals[s[i+1]] {
			result -= romanNumerals[s[i]]
		} else {
			result += romanNumerals[s[i]]
		}
	}

	return result
}
