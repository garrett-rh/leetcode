package plus

func PlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	allNines := make([]int, len(digits)+1)
	allNines[0] = 1
	return allNines
}
