package goodPairs

import "fmt"

func NumIdenticalPairs(nums []int) int {
	var pairs int
	dict := map[int]int{}

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			pairs += dict[num]
			dict[num] += 1
		} else {
			dict[num] = 1
		}
		fmt.Println(num, dict[num])
	}
	return pairs
}
