package reverseint

import "math"

func Reverse(x int) int {
	var answer int
	for x != 0 {
		answer = answer*10 + x%10
		if answer > math.MaxInt32 || answer < math.MinInt32 {
			return 0
		}

		x = x / 10
	}

	return answer
}
