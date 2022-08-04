package pivot

func pivotIndex(nums []int) int {
	var (
		sum     int
		leftsum int
	)

	for _, val := range nums {
		sum += val
	}

	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		}
		leftsum += nums[i]
	}

	return -1
}
