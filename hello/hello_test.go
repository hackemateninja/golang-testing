package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Herman")
	want := "Hello, Herman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
