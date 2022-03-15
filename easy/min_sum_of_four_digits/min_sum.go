package minSum

import (
	"sort"
	"strconv"
)

func MinimumSum(num int) int {

	numSlice := []int{}
	for num > 0 {
		numSlice = append(numSlice, num%10)
		num = num / 10
	}
	sort.Ints(numSlice)

	num1, _ := strconv.Atoi(strconv.Itoa(numSlice[0]) + strconv.Itoa(numSlice[3]))
	num2, _ := strconv.Atoi(strconv.Itoa(numSlice[1]) + strconv.Itoa(numSlice[2]))

	return num1 + num2
}
