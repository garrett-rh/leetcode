package running_sum

func RunningSum(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}
