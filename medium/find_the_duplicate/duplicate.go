package duplicates

func FindDuplicates(nums []int) int {
	var (
		slow = nums[0]
		fast = nums[nums[0]]
	)

	if len(nums) == 1 {
		return -1
	}

	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}
