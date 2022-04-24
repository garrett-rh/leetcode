package tinyUrl

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	got := Constructor()

	if reflect.TypeOf(got) != reflect.TypeOf(Codec{}) {
		t.Errorf("got type: %T, want type %T", reflect.TypeOf(got), reflect.TypeOf(Codec{}))
	}
}

func TestEncode(t *testing.T) {
	codec := Constructor()

	got := codec.encode("https://leetcode.com/problems/design-tinyurl")

	want := base64.StdEncoding.EncodeToString([]byte("https://leetcode.com/problems/design-tinyurl"))
	want = want[:6]

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDecode(t *testing.T) {
	codec := Constructor()
	got := codec.decode(codec.encode("https://leetcode.com/problems/design-tinyurl"))

	want := "https://leetcode.com/problems/design-tinyurl"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
