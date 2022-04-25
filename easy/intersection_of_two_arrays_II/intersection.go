package intersectionArrayII

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}
