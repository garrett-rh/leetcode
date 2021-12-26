package richest_customer

import "testing"

func TestMaximumWealth(t *testing.T) {
	got := MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})
	expected := 10

	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}
