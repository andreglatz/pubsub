package uuid

import (
	"testing"
)

func TestNewV4(t *testing.T) {
	got := NewV4()
	want := NewV4()

	if got == want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
