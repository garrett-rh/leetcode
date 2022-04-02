package reverse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		s    []byte
		want []byte
	}{
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Reverse string %v", test.s), func(t *testing.T) {
			got := ReverseString(test.s)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
