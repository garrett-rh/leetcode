package running_sum

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	got := RunningSum([]int{1, 2, 3, 4})
	expected := []int{1, 3, 6, 10}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
