package permutation

func BuildArray(nums []int) []int {
	var newList []int
	for i := 0; i < len(nums); i++ {
		newList = append(newList, nums[nums[i]])
	}
	return newList
}
