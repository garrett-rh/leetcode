package permutation

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	output := BuildArray([]int{0, 2, 1, 5, 3, 4})
	expected := []int{0, 1, 2, 4, 5, 3}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected %v but got %v", expected, output)
	}
}
