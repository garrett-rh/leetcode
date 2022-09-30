package reverseint

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{num: 321, want: 123},
		{num: -123, want: -321},
		{num: 120, want: 21},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Reversing integer %d", test.num), func(t *testing.T) {
			got := Reverse(test.num)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
