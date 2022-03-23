package brokenCalc

func brokenCalc(startValue int, target int) int {
	var minOps int
	for target > startValue {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		minOps++
	}

	return minOps + startValue - target
}
