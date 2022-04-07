package stoneSmash

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		stones []int
		want   int
	}{
		{stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{stones: []int{1}, want: 1},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("smashed largest two stones in %v", test.stones), func(t *testing.T) {

			got := LastStoneWeight(test.stones)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
