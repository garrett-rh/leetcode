package concat

func GetConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
