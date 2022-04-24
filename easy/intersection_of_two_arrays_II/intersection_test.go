package intersectionArrayII

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2, 2}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("intersects in %v & %v", test.nums1, test.nums2), func(t *testing.T) {
			got := Intersect(test.nums1, test.nums2)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
