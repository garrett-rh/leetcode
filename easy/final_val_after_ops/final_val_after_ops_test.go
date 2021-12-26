package after_ops

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	output := FinalValueAfterOperations([]string{"++X", "X++", "X++"})
	expected := 3

	if output != expected {
		t.Errorf("expected %v got %v", expected, output)
	}
}
