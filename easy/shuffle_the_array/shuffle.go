package shuffle

func Shuffle(nums []int, n int) []int {
	/*
		Each 'nums' has a length of 2n
		So in a list of {1, 2, 3, 4} and a n of 2
		x1 = 1, x2 = 2, y1 = 3, y2 = 4

		The end result should be x1, y1, x2, y2

		for i in nums
			if i < n
				xList.append(nums[i])
			else
				yList.append(nums[i])

		for i in range(n)
			finalList.append(xList[i])
			finalList.append(yList[i])
	*/
	var xList []int
	var yList []int
	var newList []int
	for i := 0; i < 2*n; i++ {
		if i < n {
			xList = append(xList, nums[i])
		} else {
			yList = append(yList, nums[i])
		}
	}

	for i := 0; i < n; i++ {
		newList = append(newList, xList[i])
		newList = append(newList, yList[i])
	}
	return newList
}
