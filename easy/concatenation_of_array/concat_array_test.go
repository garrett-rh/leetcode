package concat

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	output := GetConcatenation([]int{1, 2, 1})
	expected := []int{1, 2, 1, 1, 2, 1}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected %v but got %v", expected, output)
	}
}
