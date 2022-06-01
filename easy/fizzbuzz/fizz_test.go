package fizz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		num  int
		want []string
	}{
		{num: 3, want: []string{"1", "2", "Fizz"}},
		{num: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{num: 15, want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("fizzbuzz test on %d", test.num), func(t *testing.T) {
			got := FizzBuzz(test.num)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
