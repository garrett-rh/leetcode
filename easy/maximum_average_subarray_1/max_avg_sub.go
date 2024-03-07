package maximumaveragesubarray1

func findMaxAverage(nums []int, k int) float64 {
	// get original max
	var curr int
	for i := 0; i < k; i++ {
		curr += nums[i]
	}
	// save off the current max
	max := curr

	for ptr := k; ptr < len(nums); ptr++ {
		// removes the first  addition to the max, adds the latest addition
		curr = curr - nums[ptr-k] + nums[ptr]
		// if current max greater than previous max, set curr to now
		if curr > max {
			max = curr
		}
	}

	// by the time the loop exits, this will be the max
	return float64(max) / float64(k)
}
