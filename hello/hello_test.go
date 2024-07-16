package hello

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Herman", "")
		want := "Hello, Herman"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' When an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})
}
