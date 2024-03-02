package sortedsquares

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if abs(nums[left]) < abs(nums[right]) {
			ans[i] = nums[right] * nums[right]
			right--
		} else {
			ans[i] = nums[left] * nums[left]
			left++
		}
	}

	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
