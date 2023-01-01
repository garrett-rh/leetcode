package majority

func MajorityElement(nums []int) int {
	// Second solution

	dict := map[int]int{}

	for _, num := range nums {
		dict[num]++
		if dict[num] > len(nums)/2 {
			return num
		}

	}

	return -1

	// First solution
	//	dict := map[int]int{}
	//	for _, num := range nums {
	//		dict[num]++
	//		fmt.Println(dict)
	//	}
	//
	//	var largest int
	//	for key, value := range dict {
	//		if value > dict[largest] {
	//			largest = key
	//		}
	//	}
	//  return largest
}
