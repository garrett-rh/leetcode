package iteration

func TwoSum(nums []int, target int) []int {
	// Slow slow
	// var indexes []int
	// for i := 0; i < len(nums); i++ {
	// 	for idx, num := range nums {
	// 		if idx == 0 {
	// 			continue
	// 		} else if nums[i]+num == target && i != idx {
	// 			indexes = append(indexes, i)
	// 			indexes = append(indexes, idx)
	// 			break
	// 		}
	// 	}
	// 	if len(indexes) == 2 {
	// 		break
	// 	}
	// }

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
