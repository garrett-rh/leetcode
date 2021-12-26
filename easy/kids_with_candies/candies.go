package candies

func KidsWithCandies(candies []int, extraCandies int) []bool {
	largestIndex := 0
	var moreThanOthers []bool
	for idx, candy := range candies {
		if candy > candies[largestIndex] {
			largestIndex = idx
		}
	}
	for _, candy := range candies {
		if candy+extraCandies >= candies[largestIndex] {
			moreThanOthers = append(moreThanOthers, true)
		} else {
			moreThanOthers = append(moreThanOthers, false)
		}
	}
	return moreThanOthers
}
