package main

import "testing"

func TestHello(t *testing.T)  {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Alex", "Spanish")
		want := "Hola, Alex"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in french", func(t *testing.T) {
		got := Hello("", "french")
		want := " Bonjor, Alex"
		assertCorrectMessage(t, got, want)
	})
}