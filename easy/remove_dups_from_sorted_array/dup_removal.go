package removeDuplicates

func RemoveDuplicates(nums []int) int {
	var (
		slow = 0
		fast = 1
		n    = len(nums)
	)

	if n == 1 {
		return 1
	}

	for fast < n {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
