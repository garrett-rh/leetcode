package searchInsert

func SearchInsert(nums []int, target int) int {
	for idx := range nums {
		if nums[idx] >= target {
			return idx
		}
	}
	return len(nums)
}
