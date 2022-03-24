package boatLimit

import (
	"fmt"
	"testing"
)

func TestNumRescureBoats(t *testing.T) {
	testCases := []struct {
		people []int
		limit  int
		want   int
	}{
		{people: []int{1, 2}, limit: 3, want: 1},
		{people: []int{3, 2, 2, 1}, limit: 3, want: 3},
		{people: []int{3, 5, 3, 4}, limit: 5, want: 4},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Boat limit %d, people weight onboard %v", test.limit, test.people), func(t *testing.T) {
			got := NumRescueBoats(test.people, test.limit)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
