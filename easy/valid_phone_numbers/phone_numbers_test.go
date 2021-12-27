package phone_numbers

import (
	"reflect"
	"testing"
)

func TestPhoneRegex(t *testing.T) {
	got := PhoneRegex()
	want := []string{"987-123-4567", "(123) 456-7890"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
