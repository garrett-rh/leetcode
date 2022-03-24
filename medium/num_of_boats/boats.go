package boatLimit

import (
	"sort"
)

func NumRescueBoats(people []int, limit int) int {
	var (
		anchor int
		end    = len(people) - 1
	)
	sort.Sort(sort.Reverse(sort.IntSlice(people)))
	for anchor <= end {
		if people[anchor]+people[end] <= limit {
			end--
		}
		anchor++
	}
	return anchor
}
