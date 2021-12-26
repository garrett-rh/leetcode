package shuffle

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	got := Shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	expected := []int{2, 3, 5, 4, 1, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
