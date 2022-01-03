package iteration

func TwoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i, num := range nums {
		potentialMatch := target - num

		if j, found := dict[potentialMatch]; found {
			return []int{i, j}
		}
		dict[num] = i
	}
	return []int{}
}
