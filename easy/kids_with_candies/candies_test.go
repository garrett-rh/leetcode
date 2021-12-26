package candies

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	got := KidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	want := []bool{true, true, true, false, true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
