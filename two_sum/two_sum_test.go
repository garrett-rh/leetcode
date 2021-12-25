package iteration

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	output := TwoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}

	sort.Ints(output)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected %v but got %v", expected, output)
	}
}
