package defanging

import "testing"

func TestDefangIPAddr(t *testing.T) {
	got := DefangIPAddr("1.1.1.1")
	expected := "1[.]1[.]1[.]1"

	if got != expected {
		t.Errorf("got %v expected %v", got, expected)
	}
}
