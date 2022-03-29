package rotatedAndSortedSearchII

func Search(nums []int, target int) bool {
	var (
		start = 0
		end   = len(nums) - 1
	)

	if len(nums) == 0 {
		return false
	}

	for start+1 < end {
		//finding the midpoint
		mid := (end + start) / 2

		//if we are lucky enough that the midpoint is the target, return true
		if nums[mid] == target {
			return true
		}

		// if the starting point is less than our midpoint then...
		if nums[start] < nums[mid] {
			/* if the start is less than our target & the midpoint is greater tha
			the new start is the midpoint since its in array 2

			otherwise, the new end is the midpoint since its in array 1
			*/
			if nums[start] <= target && nums[mid] > target {
				end = mid
			} else {
				start = mid
			}
			/*
				if the start num is greater than our midpoint number
				and the midpoint num is less than the target while the ending number is greater than or equal to the target
				we set the start to midpoint since the target is in array 2
				otherwise the target is in array 1
			*/
		} else if nums[start] > nums[mid] {
			if nums[mid] < target && nums[end] >= target {
				start = mid
			} else {
				end = mid
			}
			/* if all fails, move the start up (can happen if the array is rotated on a duplicate entry)
			 */
		} else {
			start++
		}
	}

	/* if our start or end number is the target return true
	we can get here once the start + 1 is greater than our ending point
	*/
	if nums[start] == target || nums[end] == target {
		return true
	} else {
		// only triggers if the number is not in the list
		return false
	}
}
