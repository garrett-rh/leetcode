package rotate

func Rotate(nums []int, k int) {
	if k = k % len(nums); k != 0 {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}

}
