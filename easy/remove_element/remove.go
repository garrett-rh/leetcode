package remove

import "sort"

func RemoveElement(nums []int, val int) int {

	count := len(nums)
	for idx, num := range nums {
		if num == val {
			count--
			remove(nums, idx)
		}
	}
	sort.Ints(nums)
	return count
}

func remove(s []int, i int) {
	s[i] = 101
}
