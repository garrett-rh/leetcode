package goodPairs

func NumIdenticalPairs(nums []int) int {
	var pairs int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pairs++
			}
		}
	}
	return pairs
}
